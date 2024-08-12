package session

import (
	"log"
	"os"
)

const TempFilePath = "/var/tmp/session-monitor"

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CreateStateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating state file: %v\n", err)
		return
	}
	defer file.Close()
}
