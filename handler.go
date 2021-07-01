package main

import (
	"net/http"
	"text/template"
)

// HTTPハンドラを集めた型を定義する
type Handlers struct {
	ab *Accountbook
}

// 新しいHandlerを作成する関数を定義する
func NewHandlers(ab *Accountbook) *Handlers {
	return &Handlers{ab}
}

// ListHandlerで使用するテンプレートを定義する
var listTmpl = template.Must(template.New("list").Parse(`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8"/>
		<title>家計簿</title>
	</head>
	<body>
		<h1>家計簿</h1>
		<h2>最新{{len .}}件</h2>
		{{- if . -}}
		<table border="1">
			<tr><th>品目</th><th>値段</th></tr>
			{{- range .}}
			<tr><td>{{.Category}}</td><td>{{.Price}}円</td></tr>
			{{- end}}
		</table>
		{{- else}}
			データがありません
		{{- end}}
	</body>
</html>
`))

// 最新の入力データを表示するハンドラを定義する
func (hs *Handlers) ListHandler(w http.ResponseWriter, r *http.Request) {
	// 最新の10件を取得して、itemsに代入する
	items, err := hs.ab.GetItems(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 取得したitemsをテンプレートに埋め込む
	if err := listTmpl.Execute(w, items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
