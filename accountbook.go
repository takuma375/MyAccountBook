package main

import (
	"bufio"
	"database/sql"
	"errors"
	"os"
	"strconv"
	"strings"
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

	// 読み込み専用でファイルを開く
	file, err := os.Open(ab.fileName)
	if err != nil {
		return nil, err
	}

	// 関数終了時にファイルを閉じる
	defer file.Close()

	// ファイルからデータを読み込む
	scanner := bufio.NewScanner(file)

	// 戻り値を保管するための変数を定義
	var items []*Item

	// 読み込んだファイルを1行ずつ処理する
	for scanner.Scan() {
		// 1行ずつパースする
		var item Item
		if err := ab.parseLine(scanner.Text(), &item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	// スキャナーでエラーが発生していないか確認
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// limit件よりも少ない場合は、全件返す
	if len(items) < limit {
		return items, nil
	}

	// limit件よりも多い場合は、itemsの後方limit件だけ返す
	return items[len(items)-limit : len(items) : len(items)], nil
}

// データをパースするためのメソッドを定義
func (ab *Accountbook) parseLine(line string, item *Item) error {
	// 1行をスペースで分割する処理を追加
	splited := strings.Split(line, " ")

	// 2つに分割できなかった場合は、エラーを返す
	if len(splited) != 2 {
		return errors.New("パースに失敗しました")
	}

	// 品目名を代入
	category := splited[0]

	// 値段を代入
	price, err := strconv.Atoi(splited[1])
	if err != nil {
		return err
	}

	// 呼び出し元で指定されているポインタに値を代入する
	item.Category = category
	item.Price = price

	return nil
}
