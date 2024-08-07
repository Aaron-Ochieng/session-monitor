package session

import (
	"strconv"
	"strings"
	"time"
)

func GetDate() string {
	currentDate := time.Now()

	// Define the desired format
	format := "02 Jan 2006"

	// Format the current date
	return currentDate.Format(format)
}

func GetTodaysDate(s string) int {
	val := strings.Split(GetDate(), " ")

	date, _ := strconv.Atoi(val[0])
	return date
}
