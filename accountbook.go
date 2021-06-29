package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 家計簿の処理を行う型を新たに定義する
type Accountbook struct {
	fileName string
}

// 新しいAccountbookを生成する関数を定義する
func NewAccountBook(fileName string) *Accountbook {
	return &Accountbook{fileName: fileName}
}

// ファイルに新しいitemを追加するためのメソッドを定義する
func (ab *Accountbook) AddItem(item *Item) error {
	// 追記モードでファイルを開く
	file, err := os.OpenFile(ab.fileName, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 「品目 値段」の形式でファイルに出力する
	if _, err := fmt.Fprintln(file, item.Category, item.Price); err != nil {
		log.Fatal(err)
	}

	// ファイルを閉じる
	file.Close()

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
		// TODO: データのパースを行う関数を別途定義
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
