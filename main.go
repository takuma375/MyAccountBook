package main

import "fmt"

// 値段と品目を一緒に扱うためにItemという構造体の型を定義する
type Item struct {
	Category string
	Price    int
}

func main() {

	// inputItem()を呼び出し、結果をitemという変数に代入する
	item := inputItem()

	// "==========="と出力して改行する
	fmt.Println("===========")

	// 品目に「コーヒー」、値段に「100」と入力した場合に
	// 「コーヒーに100円使いました」と表示する
	fmt.Printf("%sに%d円使いました\n", item.Category, item.Price)

	// 「===========」と出力して改行する
	fmt.Println("===========")

}

// データの入力を行う関数を定義する
func inputItem() Item {
	// 入力された値を仮保管するItem型の変数を定義
	var item Item

	// "品目>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	// "値段>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return item
}
