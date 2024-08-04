package session

import (
	"errors"
	"net"
)

func GetMacAddress() (string, error) {
	addr, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, val := range addr {
		if val.Flags&net.FlagUp == 0 || val.Flags&net.FlagLoopback != 0 {
			continue
		}

		mac_addr := val.HardwareAddr.String()
		if mac_addr != "" {
			return mac_addr, nil
		}
	}
	return "", errors.New("mac addr not found")
}
