package main

// 家計簿の処理を行う型を新たに定義する
type Accountbook struct {
	fileName string
}

// 新しいAccountbookを生成する関数を定義する
func NewAccountBook(fileName string) *Accountbook {
	return &Accountbook{fileName: fileName}
}
