package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:root@tcp(db:3306)/gotxdbsample?parseTime=True&loc=Local"

func init() {
	txdb.Register(
		"txdbsample", // 今回使用するドライバにつける任意のドライバ名
		"mysql",      // 今回使用するドライバ名
		dsn,          // DSN名
	)
}

func main() {
	db, err := sql.Open(
		"txdbsample", // txdb.Registerで設定した任意のドライバ名
		dsn,          // DSN名
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// exp) INSERT
	result, err := db.Exec(`INSERT INTO users(id, username) VALUES(4, "テストユーザー")`)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %v\n", id) // => ID: 4

	// exp) SELECT
	var username string

	if err := db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&username); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UserName: %v\n", username) // => UserName: テストユーザー
}
