package session

func from(logs []LoginInfo) (res int) {
	day := CurrentDate()
	for i := len(logs) - 1; i >= 0; i-- {
		res += 1
		if logs[i].Date == day {
			break
		}
	}
	return res
}

