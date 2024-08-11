package session

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func connection() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "./records.db")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return db, nil
}

