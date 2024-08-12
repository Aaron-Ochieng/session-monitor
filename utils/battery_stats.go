package session

import (
	"fmt"
	"os/exec"
	"strings"
)

type BatStats struct {
	State       string
	TimetoEmpty string
	Percentage  string
}

func BatteryStatus() (stats BatStats) {
	cmd := exec.Command("bash", "-c", `upower -i $(upower -e | grep BAT) | grep --color=never -E "state|to\ full|to\ empty|percentage"`)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	res := strings.Split(string(output), "\n")

	if len(res) == 4 {
		stats.State = trimspace(res[0])[1]
		stats.TimetoEmpty = trimspace(res[1])[3] + " " + trimspace(res[1])[4]
		stats.Percentage = trimspace(res[2])[1]

	} else if len(res) == 2 {
		stats.State = trimspace(res[0])[1]
		stats.TimetoEmpty = trimspace(res[1])[3] + " " + trimspace(res[1])[4]
		stats.Percentage = trimspace(res[2])[1]

	} else if len(res) == 3 {
		stats.State = trimspace(res[0])[1]
		stats.Percentage = trimspace(res[1])[1]
	}
	return stats
}
