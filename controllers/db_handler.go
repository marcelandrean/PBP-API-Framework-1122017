package controllers

import (
	"database/sql"
	"echo-rest/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
