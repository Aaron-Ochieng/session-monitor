package main

import (
	"fmt"

	session "session/utils"
)

func main() {
	// check if its first run
	is_first_run := !session.FileExists(session.TempFilePath)
	if is_first_run {
		session.InitDbTables()
		session.CreateStateFile(session.TempFilePath)
	}
	logs, _ := session.UnixLog()
	session.InsertUsername(logs)
	session.InsertLog(is_first_run)
	fmt.Println(session.BatteryStatus())
}
