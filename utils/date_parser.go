package session

import (
	"time"
)

func GetDate() string {
	currentDate := time.Now()

	// Define the desired format
	format := "02 Jan 2006"

	// Format the current date
	return currentDate.Format(format)
}