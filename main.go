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

}
