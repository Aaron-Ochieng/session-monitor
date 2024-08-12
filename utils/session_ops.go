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

func LogsToInsert(logs []LoginInfo, from int, last_log_date string) []LoginInfo {
	to := 0
	for i := len(logs) - 1; i >= 0; i-- {
		to += 1
		if logs[i].Date == last_log_date {
			break
		}
	}
	return logs[from:to]
}

