package dao

import (
	"context"
	"database/sql"
	"fmt"
	exchange "go-extensive-client-server-api/server/models"
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

func SaveRequest(cambioQuotation exchange.Quotation) error {
	db := initializeDB()

	defer db.Close()
	return insert(db, cambioQuotation)
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

func FindAllQuotations(db *sql.DB) []exchange.Quotation {
	row, err := db.Query("SELECT code, name, high, low, var_bid, pct_change, bid, ask,	timestamp, create_date FROM request_quotation ORDER BY id")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer row.Close()
	var result []exchange.Quotation
	for row.Next() {
		var cambioQuotation exchange.Quotation
		row.Scan(&cambioQuotation.USDBRL.Code, &cambioQuotation.USDBRL.Name, &cambioQuotation.USDBRL.High,
			&cambioQuotation.USDBRL.Low, &cambioQuotation.USDBRL.VarBid, &cambioQuotation.USDBRL.PctChange,
			&cambioQuotation.USDBRL.Bid, &cambioQuotation.USDBRL.Ask, &cambioQuotation.USDBRL.Timestamp,
			&cambioQuotation.USDBRL.CreateDate)
		result = append(result, cambioQuotation)
		fmt.Printf("\nCotação: %v, Moeda: %v, Data: %v", cambioQuotation.USDBRL.Bid, cambioQuotation.USDBRL.Name, cambioQuotation.USDBRL.CreateDate)
	}
	return result
}

func insert(db *sql.DB, cambioQuotation exchange.Quotation) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	sqlInsert := `INSERT INTO request_quotation(code, name, high, low, var_bid, pct_change, bid, ask,	timestamp, create_date) 
	                                           VALUES (   ?,    ?,    ?,   ?,       ?,          ?,   ?,   ?,          ?,           ?)`

	statement, err := db.PrepareContext(ctx, sqlInsert)
	if err != nil {
		return err
	}
	_, err = statement.Exec(cambioQuotation.USDBRL.Code, cambioQuotation.USDBRL.Name, cambioQuotation.USDBRL.High,
		cambioQuotation.USDBRL.Low, cambioQuotation.USDBRL.VarBid, cambioQuotation.USDBRL.PctChange,
		cambioQuotation.USDBRL.Bid, cambioQuotation.USDBRL.Ask, cambioQuotation.USDBRL.Timestamp,
		cambioQuotation.USDBRL.CreateDate)

	return err
}
