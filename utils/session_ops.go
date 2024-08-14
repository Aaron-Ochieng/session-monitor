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

func FirstLogs(logs []LoginInfo) []LoginInfo {
	todays_date := CurrentDate()
	from := 0
	for i := len(logs) - 1; i >= 0; i-- {
		from += 1
		if logs[i].Date == todays_date {
			break
		}
	}
	return logs[len(logs)-from+1:]
}

func InsertLog(is_first_insert bool) {
	// check if its the first time to insert the log
	logs, _ := UnixLog()

	if is_first_insert {
		first_logs := FirstLogs(logs)
		InsertLogs(first_logs, InsertLogDate)
		return
	}
	logs_to_insert := LogsToInsert(logs, from(logs), LastLogDate())
	InsertLogs(logs_to_insert, UpdateLogDate)
}
