package session

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func PreviousDate() string {
	currentTime := time.Now()
	return currentTime.AddDate(0, 0, -1).Format("2 Jan 2004")
}

func CurrentDate() string {
	return time.Now().Format("2 Jan 2006")
}

func LogRange(arrLogs []LoginInfo) (logs []LoginInfo) {
	for _, logInfo := range arrLogs {
		if logInfo.Date == "7 Aug 2024" {
			fmt.Println("OK")
		}
	}
	return logs
}

func convertToHours(timeStr string) (float64, error) {
	var days, hours, minutes int
	var err error

	// Split the time string into day and time parts
	parts := strings.Split(timeStr, "+")

	// If there is a day part, parse it
	if len(parts) == 2 {
		days, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("invalid day format: %v", err)
		}
		timeStr = parts[1]
	}

	// Split the time part into hours and minutes
	timeParts := strings.Split(timeStr, ":")
	if len(timeParts) != 2 {
		return 0, fmt.Errorf("invalid time format")
	}

	// Parse hours
	hours, err = strconv.Atoi(timeParts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hour format: %v", err)
	}

	// Parse minutes
	minutes, err = strconv.Atoi(timeParts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minute format: %v", err)
	}

	// Calculate total hours
	totalHours := float64(days*24) + float64(hours) + float64(minutes)/60.0
	return totalHours, nil
}