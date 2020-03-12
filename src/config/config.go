package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlDb() (db *sql.DB, err error) { //con trỏ sql.db
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1605"
	dbName := "golangweb"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
