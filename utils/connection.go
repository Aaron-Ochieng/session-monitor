package session

import (
	"database/sql"
	"fmt"
)

func Connection(credentials string) *sql.DB {
	db, err := sql.Open("postgres", credentials)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()
	return db
}
