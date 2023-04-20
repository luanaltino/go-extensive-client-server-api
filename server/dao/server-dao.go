package dao

import (
	"context"
	"database/sql"
	"fmt"
	"go-extensive-client-server-api/server/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func initializeDB() *sql.DB {

	db, err := sql.Open("sqlite3", "./databasesqlite3")

	if err != nil {
		panic(err)
	}

	createTable(db)

	return db
}

func SaveRequest(cambioQuotation models.CambioQuotation) {
	db := initializeDB()

	insertStudent(db, cambioQuotation)

	findAllQuotations(db)

	defer db.Close()
}

func createTable(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS request_quotation (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"high" TEXT,
		"low" TEXT,
		"var_bid" TEXT,
		"pct_change" TEXT,
		"bid" TEXT,
		"ask" TEXT,
		"timestamp" TEXT,
		"create_date" TEXT
	  );`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt.Exec()
	defer stmt.Close()
}

func findAllQuotations(db *sql.DB) {
	row, err := db.Query("SELECT code, name, high, low, var_bid, pct_change, bid, ask,	timestamp, create_date FROM request_quotation ORDER BY id")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer row.Close()
	for row.Next() {
		var cambioQuotation models.CambioQuotation
		row.Scan(&cambioQuotation.USDBRL.Code, &cambioQuotation.USDBRL.Name, &cambioQuotation.USDBRL.High,
			&cambioQuotation.USDBRL.Low, &cambioQuotation.USDBRL.VarBid, &cambioQuotation.USDBRL.PctChange,
			&cambioQuotation.USDBRL.Bid, &cambioQuotation.USDBRL.Ask, &cambioQuotation.USDBRL.Timestamp,
			&cambioQuotation.USDBRL.CreateDate)
		fmt.Printf("\nCotação: %v, Moeda: %v, Data: %v", cambioQuotation.USDBRL.Bid, cambioQuotation.USDBRL.Name, cambioQuotation.USDBRL.CreateDate)
	}
}

func insertStudent(db *sql.DB, cambioQuotation models.CambioQuotation) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	insertStudentSQL := `INSERT INTO request_quotation(code, name, high, low, var_bid, pct_change, bid, ask,	timestamp, create_date) 
	                                           VALUES (   ?,    ?,    ?,   ?,       ?,          ?,   ?,   ?,          ?,           ?)`

	statement, err := db.PrepareContext(ctx, insertStudentSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = statement.Exec(cambioQuotation.USDBRL.Code, cambioQuotation.USDBRL.Name, cambioQuotation.USDBRL.High,
		cambioQuotation.USDBRL.Low, cambioQuotation.USDBRL.VarBid, cambioQuotation.USDBRL.PctChange,
		cambioQuotation.USDBRL.Bid, cambioQuotation.USDBRL.Ask, cambioQuotation.USDBRL.Timestamp,
		cambioQuotation.USDBRL.CreateDate)

	if err != nil {
		fmt.Println(err.Error())
	}
}
