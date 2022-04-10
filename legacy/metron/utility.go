package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DbHost             = "localhost"
	DbPort             = "3306"
	DbUser             = "root"
	DbPass             = "mypassword"
	DbName             = "tiktok"
	DbConnectionFormat = "%s:%s@tcp(%s:%s)/%s"
)

func JSON(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonStr)
}

func NewDbConn() (*sql.DB, error) {
	return sql.Open("mysql", database)
}

func NewCustomDbConn(host, port, user, pass, name string) (*sql.DB, error) {
	connection := fmt.Sprintf(DbConnectionFormat,
		user,
		pass,
		host,
		port,
		name,
	)

	return sql.Open("mysql", connection)
}
