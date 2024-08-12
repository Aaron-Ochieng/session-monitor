package session

type User struct {
	ID       int64  `pg:"id,pk,autoincrement"`
	Username string `pg:"username,unique,notnull"`
}

// UserLog represents the userlogs table.
type UserLog struct {
	ID         int64   `pg:"id,pk,autoincrement"`
	MacAddress string  `pg:"macAddress,notnull"`
	UserId     int64   `pg:"userId,notnull"`
	Date       string  `pg:"date"`
	LoginTime  string  `pg:"loginTime,notnull"`
	LogoutTime string  `pg:"logoutTime"`
	Uptime     float64 `pg:"hours"`
	User       *User   `pg:"rel:has-one,join_fk:userId"`
}

type LoginInfo struct {
	Username   string
	Date       string
	LoginTime  string
	LogoutTime string
	DeviceId   string
	Uptime     float64
}
