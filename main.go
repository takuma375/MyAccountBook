package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// accountbook.goを用いて処理を分割する

func main() {

	// NewAccountBookを使用して、Accountbookを作成する
	ab := NewAccountBook("accountbook.txt")

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

// 入力されたデータの一覧表示を行う関数
// データはファイルから直接参照する
func showItems() error {

	// "accountbook.txt"を読み込み専用で開く
	file, err := os.Open("accountbook.txt")
	if err != nil {
		log.Fatal(err)
	}

	// "==========="と出力して改行する
	fmt.Println("===========")

	// ファイルからデータを読み込む
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// 1行分のデータを取り出す
		line := scanner.Text()

		splited := strings.Split(line, " ")
		if len(splited) != 2 {
			return fmt.Errorf("パースに失敗しました")
		}

		// categoryを取り出す
		category := splited[0]
		// priceを取り出す。sting型からint型に変換することを忘れない
		price, err := strconv.Atoi(splited[1])
		if err != nil {
			log.Fatal(err)
		}

		// 「コーヒー:100円」のように表示する
		fmt.Printf("%s:%d円\n", category, price)

	}
	// 「===========」と出力して改行する
	fmt.Println("===========")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
