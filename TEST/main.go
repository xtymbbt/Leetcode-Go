package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // init mysql driver
	log "github.com/sirupsen/logrus"
)

func main() {
	// dsn := "user:password@tcp(url:3306)/db_name"
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("open  in 127.0.0.1:3306 failed,\nerror is: %v\n", err)
		return
	}
	db.SetMaxIdleConns(16) // 最大闲置连接数
	db.SetMaxOpenConns(100) // 最大连接数
	if err = db.Ping(); err != nil {
		log.Errorf("open  in 127.0.0.1:3306 failed,\nerror is: %v\n", err)
		return
	}

	tx, err := db.Begin()

	if err != nil {
		log.Errorf("begin tx err, error is: %v\n", err)
	}
	rows, err := tx.Query("use test")
	if err != nil {
		log.Errorf("begin tx err, error is: %v\n", err)
	}
	err = rows.Close()
	if err != nil {
		log.Errorf("rows close err, error is: %v\n", err)
	}
	rows, err = tx.Query("insert into test1 values (7)")
	if err != nil {
		log.Errorf("sql query err, error is: %v\n", err)
	}
	err = rows.Close()
	if err != nil {
		log.Errorf("rows close err, error is: %v\n", err)
	}
	rows, err = tx.Query("use test1")
	if err != nil {
		log.Errorf("sql query err, error is: %v\n", err)
	}
	err = rows.Close()
	if err != nil {
		log.Errorf("rows close err, error is: %v\n", err)
	}
	rows, err = tx.Query("insert into test1 values (7)")
	if err != nil {
		log.Errorf("sql query err, error is: %v\n", err)
	}
	err = rows.Close()
	if err != nil {
		log.Errorf("rows close err, error is: %v\n", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Errorf("tx commit err, error is: %v\n", err)
	}

}