package session

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
