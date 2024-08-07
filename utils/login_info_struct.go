package session

type FDate struct {
	Day   string
	Month string
	Year  int
}

type LoginInfo struct {
	Username   string
	Date       FDate
	LoginTime  string
	LogoutTime string
	DeviceId   string
}
