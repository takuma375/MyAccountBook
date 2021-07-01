package main

// HTTPハンドラを集めた型を定義する
type Handlers struct {
	ab *Accountbook
}

// 新しいHandlerを作成する関数を定義する
func NewHandlers(ab *Accountbook) *Handlers {
	return &Handlers{ab}
}
