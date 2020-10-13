package main

import (
	"context"
	"database/sql"
	"errors"

	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var server = "127.0.0.1"
var port = 64077
var user = "testuser"
var password = "123"
var database = "loger"
var sizelimit int64 = 150

func connectsql(Tablname string, Colonsname string, Log string) error {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)

	if err != nil {
		return err
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	// Create pole
	create, err := Create(Tablname, Colonsname, Log)
	if err != nil {
		fmt.Println("Error creating: ", err)
		err = nil
		return err
	}

	fmt.Println(create)
	return err
}

// Create inserts an pole record
func Create(Tablname string, Colonsname string, Log string) (string, error) {
	ctx := context.Background()
	var err error
	if db == nil {
		err = errors.New("Createpole: db is null")
		return "Ошибка", err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return "Ошибка", err
	}
	tsql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ('%s')`, Tablname, Colonsname, Log)
	stmt, err := db.Query(tsql)
	if err != nil {
		return "Ошибка", err
	}
	defer stmt.Close()

	var new string
	return new, nil
}
