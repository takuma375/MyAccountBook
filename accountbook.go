package main

import (
	"database/sql"
)

// 値段と品目を一緒に扱うためにItemという構造体の型を定義する
// 今回からRDBを使用するため、フィールドにIDを定義する
type Item struct {
	ID       int
	Category string
	Price    int
}

// 家計簿の処理を行う型を定義する。データの管理にはデータベースを使用する
type Accountbook struct {
	db *sql.DB
}

// 新しいAccountbookを生成する関数を定義。
func NewAccountBook(db *sql.DB) *Accountbook {
	return &Accountbook{db: db}
}

// データ管理のためのテーブルを作成するメソッドを定義
func (ab *Accountbook) CreateTable() error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS items(
		id        INTEGER PRIMARY KEY,
		category  TEXT NOT NULL,
		price     INTEGER NOT NULL
	);`

	_, err := ab.db.Exec(sqlStr)
	if err != nil {
		return err
	}

	return nil
}

// ファイルに新しいitemを追加するためのメソッドを定義する。RDBを使った仕様に変更する。
func (ab *Accountbook) AddItem(item *Item) error {
	// SQLのInsert文を使って、データベースに値を保存する
	// ?の部分にcategoryやpriceを代入できるようにする。
	const sqlStr = `INSERT INTO items(category, price) VALUES (?,?);`
	_, err := ab.db.Exec(sqlStr, item.Category, item.Price)
	if err != nil {
		return err
	}
	// 成功終了したことを伝えるため、nilを返す
	return nil
}

// 最近追加したものを最大limit件だけItemを取得するメソッドを定義する。もしエラーが発生したら第2戻り値で返す
func (ab *Accountbook) GetItems(limit int) ([]*Item, error) {

	const sqlStr = `SELECT * FROM items ORDER BY id DESC LIMIT ?`
	rows, err := ab.db.Query(sqlStr, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close() // 関数終了時にCloseする

	var items []*Item

	// rows.Next()を使用して1行ずつ取得した行を見る
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Category, &item.Price); err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// 品目ごとに集計するメソッドを定義する
func (ab *Accountbook) GetSummaries() ([]*Summary, error) {
	const sqlStr = `
		SELECT
			category,
			COUNT(1) as count,
			SUM(price) as sum
		FROM
			items
		GROUP BY
			category`
	rows, err := ab.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 関数終了時にCloseする

	var summaries []*Summary
	for rows.Next() {
		var s Summary
		err := rows.Scan(&s.Category, &s.Count, &s.Sum)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return summaries, nil

}

// 集計したデータを扱う構造体Summaryを定義する
type Summary struct {
	Category string
	Count    int
	Sum      int
}

// 平均を取得するメソッドを定義する
func (s *Summary) Avg() float64 {
	// Countが0だと除算になるため、そのまま0を返す
	if s.Count == 0 {
		return 0
	}
	// float64にキャストして返す
	return float64(s.Sum) / float64(s.Count)
}
