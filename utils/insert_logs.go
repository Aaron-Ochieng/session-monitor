package session

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/joho/godotenv"
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
		user := &User{
			Username: log.Username,
		}

		_, err := db.Model(user).
			OnConflict("(username) DO UPDATE").
			Set("username = EXCLUDED.username").
			Insert()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func InsertLogs(logs []LoginInfo, logsDate func()) {
	db, err := connection()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	macAddress, _ := GetMacAddress()

	for _, log := range logs {
		user := &User{Username: log.Username}

		// Fetch or create the user by username
		_, err := db.Model(user).
			Where("username = ?", log.Username).
			OnConflict("DO NOTHING").
			SelectOrInsert()
		if err != nil {
			fmt.Println("Error selecting or inserting user:", err)
			continue
		}

		userLog := &UserLog{
			MacAddress: macAddress,
			UserId:     user.ID, // Automatically get userId from the User struct
			LoginTime:  log.LoginTime,
			LogoutTime: log.LogoutTime,
			Date:       log.Date,
			Uptime:     log.Uptime,
		}

		_, err = db.Model(userLog).Insert()
		if err != nil {
			fmt.Println("Error inserting log:", err)
		}
	}

	logsDate()
}

func LastLogDate() (res string) {
	db, err := connection()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	macAddress, err := GetMacAddress()
	if err != nil {
		fmt.Println("Error getting Mac address:", err)
		return ""
	}

	lastInsertDate := &LastInsertDate{}
	err = db.Model(lastInsertDate).
		Where("mac_address = ?", macAddress).
		Select()
	if err != nil {
		fmt.Println("Error fetching last log date:", err)
		return ""
	}

	return lastInsertDate.Date
}

func InsertLogDate() {
	db, err := connection()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	logDate := CurrentDate()
	macAddress, _ := GetMacAddress()

	lastInsertDate := &LastInsertDate{
		MacAddress: macAddress,
		Date:       logDate,
	}

	_, err = db.Model(lastInsertDate).
		Insert()
	if err != nil {
		log.Fatalln("Error inserting or updating last log date:", err)
	}
}

func UpdateLogDate() {
	db, err := connection()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	logDate := CurrentDate()
	macAddress, _ := GetMacAddress()

	_, err = db.Model(&LastInsertDate{}).
		Set("date = ?", logDate).
		Where("mac_address = ?", macAddress).
		Update()
	if err != nil {
		log.Fatalln("Error updating last log date:", err)
	}
}
