package session

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func InitDbTables() error {
	db, err := connection()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users(
			id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(20) NOT NULL UNIQUE
			
		);
		CREATE TABLE IF NOT EXISTS userlogs(
      		id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			macAddress VARCHAR(20) NOT NULL,
			userId INTEGER NOT NULL,
			date VARCHAR(11),
			loginTime  VARCHAR(12) NOT NULL,
			logoutTime VARCHAR(12),
			hours REAL,
			FOREIGN KEY (userId) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS lastInsertDate(
			macAddress VARCHAR(20) NOT NULL PRIMARY KEY,
			date VARCHAR(11)
		);      
		`,
	)
	if err != nil {
		return err
	}
	return nil
}

func InsertUsername(logs []LoginInfo) {
	db, err := connection()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	for _, log := range logs {
		_, err = db.ExecContext(
			context.Background(),
			`INSERT INTO users (username) VALUES (?) ON 
			CONFLICT (username) DO UPDATE SET username = excluded.username`, log.Username,
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func userId(username string) (id int) {
	db, _ = connection()
	err := db.QueryRowContext(
		context.Background(),
		`SELECT id FROM users WHERE username = ? ;`, username,
	).Scan(&id)
	if err != nil {
		log.Fatalln(err)
	}
	return id
}

func InsertLogs(logs []LoginInfo, logsDate func()) {
	db, _ = connection()
	macAddress, _ := GetMacAddress()
	for _, log := range logs {
		userId := userId(log.Username)
		_, err := db.ExecContext(
			context.Background(),
			`INSERT INTO userlogs
      			(macAddress,userId,loginTime,logoutTime,date,hours) 
      			VALUES (?,?,?,?,?,?,?);`,
			macAddress,
			userId,
			log.LoginTime,
			log.LogoutTime,
			log.Date,
			log.Uptime,
		)
		if err != nil {
			fmt.Println(err)
		}

	}
	logsDate()
	db.Close()
}

func LastLogDate() (res string) {
	db, _ := connection()
	macAddress, _ := GetMacAddress()

	db.QueryRowContext(context.Background(),
		`SELECT date FROM lastInsertDate WHERE macAddress = ?`,
		macAddress).Scan(&res)
	return res
}

func InsertLogDate() {
	db, _ := connection()
	log_date := CurrentDate()
	macAddress, _ := GetMacAddress()

	_, err := db.ExecContext(context.Background(),
		`INSERT INTO lastInsertDate (date,macAddress) VALUES (?,?)`,
		log_date, macAddress,
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateLogDate() {
	db, _ := connection()
	log_date := CurrentDate()
	macAddress, _ := GetMacAddress()

	_, err := db.ExecContext(context.Background(),
		`UPDATE lastInsertDate SET date = ? WHERE macAddress = ?`,
		log_date, macAddress,
	)
	if err != nil {
		log.Fatalln(err)
	}
}
