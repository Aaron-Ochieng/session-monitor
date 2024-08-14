package session

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

var db *pg.DB

func connection() (db *pg.DB, err error) {
	credentials := DatabaseCredentials{}
	// Load environment variables from a .env file if it exists
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using system environment variables instead")
	}

	// Retrieve connection details from environment variables
	credentials.DB_USER = os.Getenv("DB_USER")
	credentials.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	credentials.DB_NAME = os.Getenv("DB_NAME")
	credentials.DB_ADDR = os.Getenv("DB_ADDR")

	db = pg.Connect(&pg.Options{
		User:     credentials.DB_USER,
		Password: credentials.DB_PASSWORD,
		Database: credentials.DB_NAME,
		Addr:     credentials.DB_ADDR,
	})

	// Test the connection.
	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDbTables() error {
	db, err := connection()
	if err != nil {
		return err
	}
	defer db.Close()

	models := []interface{}{
		(*User)(nil),
		(*UserLog)(nil),
		(*LastInsertDate)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
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
      			VALUES (?,?,?,?,?,?);`,
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
