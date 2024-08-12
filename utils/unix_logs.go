package session

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func trimspace(s string) (res []string) {
	temp := ""
	for _, char := range s {
		if char != ' ' {
			temp += string(char)
		} else if char == ' ' && temp != "" {
			res = append(res, temp)
			temp = ""
		}
	}
	if temp != "" {
		res = append(res, temp)
	}
	return res
}

func UnixLog() (logininfo []LoginInfo, err error) {
	command, err := exec.Command("last", "-f", "/var/log/wtmp").Output()
	if err != nil {
		return nil, err
	}
	mac_address, err := GetMacAddress()
	if err != nil {
		return nil, err
	}
	res := strings.Split(string(command), "\n")
	// Remove the 2 extra space and wtmp information of startuptime
	res = res[:len(res)-3]

	for _, val := range res {
		info := trimspace(val)
		if info[0] == "bocal" {
			continue
		}
		hours, _ := convertToHours(info[9][1 : len(info[9])-1])
		temp := LoginInfo{}
		temp.Username = info[0]
		temp.Date = info[5] + " " + info[4] + " " + strconv.Itoa(time.Now().Year())
		temp.LoginTime = info[6]
		temp.LogoutTime = info[8]
		temp.Uptime = hours
		if info[8] == "logged" {
			temp.DeviceId = mac_address
		}
		if info[0] != "reboot" {
			logininfo = append(logininfo, temp)
		}
	}
	return logininfo, nil
}
