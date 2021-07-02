package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

	// HandlersをNewHandlersを使って作成する
	hs := NewHandlers(ab)

	// ハンドラを登録する
	http.HandleFunc("/", hs.ListHandler)
	// saveHandlerを登録する
	http.HandleFunc("/save", hs.SaveHandler)

	fmt.Println("http://localhost:8080 で起動中...")

	// HTTPサーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
