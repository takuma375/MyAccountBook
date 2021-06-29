package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		return err
	}

	// 「品目 値段」の形式でファイルに出力する
	if _, err := fmt.Fprintln(file, item.Category, item.Price); err != nil {
		return err
	}

	// ファイルを閉じる
	if err := file.Close(); err != nil {
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
