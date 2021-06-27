package main

import (
	"fmt"
	"log"
	"os"
)

// 値段と品目を一緒に扱うためにItemという構造体の型を定義する
type Item struct {
	Category string
	Price    int
}

// データのファイルへの保存機能を実装する

func main() {

	// "accountbook.txt"という名前のファイルを書き込み用で開く
	file, err := os.Create("accountbook.txt")
	if err != nil {
		// エラーを出力して終了
		log.Fatal(err)
	}

	// 入力するデータの件数を指定してもらうため、変数の定義と代入を行う
	var n int
	fmt.Print("何件入力しますか？>")
	fmt.Scan(&n)

	// inputItem()を呼び出し、ファイルに入力を保存する
	for i := 0; i < n; i++ {
		if err := inputItem(file); err != nil {
			log.Fatal(err)
		}
	}

	// ファイルを閉じる
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	// showItems()を呼び出し、データの一覧表示をする
	// showItems(items)

}

// データの入力を行う関数を定義する
// データの保存はテキストファイルに行うため、*os.File型の引数を受け取り、エラー処理のためのerror型の返り値を返すように変更
func inputItem(file *os.File) error {
	// 入力された値を仮保管するItem型の変数を定義
	var item Item

	// "品目>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	// "値段>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	// ファイルに「品目 値段」のように書き出す
	line := fmt.Sprintf("%s %d\n", item.Category, item.Price)
	if _, err := file.WriteString(line); err != nil {
		log.Fatal(err)
	}

	// 何もエラーが起こらなかったことを表すnilを返す
	return nil
}

// 入力されたデータの一覧表示を行う関数を新たに作成する
func showItems(items []Item) {

	// "==========="と出力して改行する
	fmt.Println("===========")

	// itemsの長さだけ、for文を回し、データを一覧表示する。
	// 「コーヒー:120円」のように表示する。
	for i := 0; i < len(items); i++ {
		fmt.Printf("%s:%d円\n", items[i].Category, items[i].Price)
	}

	// 「===========」と出力して改行する
	fmt.Println("===========")
}
