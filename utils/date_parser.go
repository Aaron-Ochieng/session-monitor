package session

import (
	"fmt"
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
