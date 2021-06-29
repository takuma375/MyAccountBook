package main

import (
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
