package monitor

import (
	"database/sql"
	"fmt"
)

func connection(credentials string) *sql.DB {
	db, err := sql.Open("postgres", credentials)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
