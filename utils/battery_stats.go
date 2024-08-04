package session

import (
	"fmt"
	"os"
	"strconv"
)

func BatteryStatus() (int, error) {
	filepath := "/sys/class/power_supply/BAT0/capacity"
	battery_capacity, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	val, _ := strconv.Atoi(string(battery_capacity))
	return val, nil
}
