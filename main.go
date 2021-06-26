package main

import "fmt"

func main() {
	// 品目を入れる変数を定義する
	var category string

	// 値段を入れる変数を定義する
	var price int

	// "品目>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("品目>")
	fmt.Scan(&category)

	// "値段>"と表示し、入力した結果を品目を入れる変数に代入する
	fmt.Print("値段>")
	fmt.Scan(&price)

	// "==========="と出力して改行する
	fmt.Println("===========")

	// 品目に「コーヒー」、値段に「100」と入力した場合に
	// 「コーヒーに100円使いました」と表示する
	fmt.Printf("%sに%d円使いました\n", category, price)

	// 「===========」と出力して改行する
	fmt.Println("===========")

}
