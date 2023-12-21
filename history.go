package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName     = "mysql"
	dataSourceName = "root:123456@tcp(127.0.0.1:3306)/module"
)

func (g *Gobang) saveBoard() {
	// 连接数据库
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 执行插入操作
	insertStmt, err := db.Prepare("INSERT INTO board (board_info) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec("value1") // 执行插入操作
	if err != nil {
		panic(err.Error())
	}
}
