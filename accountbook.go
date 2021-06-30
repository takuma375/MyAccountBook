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

}
