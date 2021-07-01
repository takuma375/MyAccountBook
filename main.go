package main

import (
	"database/sql"
	"fmt"
	"os"

	//SQLiteのドライバを使うために "github.com/tenntenn/sqlite"をインポートする
	"github.com/tenntenn/sqlite"
)

// データベースによるデータの保管

func main() {

	// データベースへの接続処理
	db, err := sql.Open(sqlite.DriverName, "accountbook.db")
	if err != nil {
		// 標準エラー出力にエラーメッセージを出力する
		fmt.Fprintln(os.Stderr, "エラー", err)
		// ステータスコード1で終了
		os.Exit(1)
	}

	// NewAccountBookを使用して、Accountbookを作成する
	ab := NewAccountBook(db)

	// テーブルを作成する処理を追加
	if err := ab.CreateTable(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー", err)
		os.Exit(1)
	}

	// 以下のループにラベルを付ける
LOOP:
	for {
		// モードを選択して実行できるようにする
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]集計 [4]終了")
		fmt.Print("> ")
		fmt.Scan(&mode)

		// modeによって処理を変える
		switch mode {
		case 1: // 入力
			// データの入力件数を受け取る
			var n int
			fmt.Print("何件入力しますか? ")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				err := ab.AddItem(inputItem())
				if err != nil {
					fmt.Fprintln(os.Stderr, "エラー", err)
					// LOOPという名前のついたforから抜け出す
					break LOOP
				}
			}
		case 2: //最新10件
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー", err)
				break LOOP
			}
			showItems(items)
		case 3: // 集計
			summaries, err := ab.GetSummaries()
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー", err)
				break LOOP
			}
			showSummary(summaries)
		case 4: // 終了
			fmt.Println("終了します")
			os.Exit(0)
		}
	}

}

// データの入力を行う関数を定義する。データの保存処理などはAddItem関数が担うため、
// ここでは入力の受け渡しを行う
func inputItem() *Item {
	// 入力された値を仮保管するItem型の変数を定義
	var item Item

	// "品目>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	// "値段>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("値段>")
	fmt.Scan(&item.Price)
	fmt.Print("\n")

	// 入力された結果を返す
	return &item
}

// 入力されたデータの一覧表示を行う関数
// accountbook.go内で処理されたデータを受け取り、出力する
func showItems(items []*Item) {

	// "==========="と出力して改行する
	fmt.Println("===========")

	for _, item := range items {
		// itemsの要素を1つずつ取り出してitemに入れて繰り返す
		fmt.Printf("[%04d]%s:%d円\n", item.ID, item.Category, item.Price)
	}

	// 「===========」と出力して改行する
	fmt.Println("===========")

}

// 集計を出力する関数を定義する
func showSummary(summaries []*Summary) {
	fmt.Println("===========")
	// タブ区切りで「品目 個数 合計 平均」を出力
	fmt.Printf("品目\t個数\t合計\t平均\n")
	for _, s := range summaries {
		fmt.Printf("%s\t%d\t%d円\t%.2f円\n", s.Category, s.Count, s.Sum, s.Avg())
	}
	fmt.Println("===========")
}
