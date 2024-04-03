package myconverts

import "time"

// conver string to time with format : 2006-01-02 15:04:05
func StrToTimeWFormat(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

// conver time to string with format : 2006-01-02 15:04:05
func TimeToStrWFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// conver string to time with format : 2006-01-02
func StrToTimeWFormatDate(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

// conver time to string with format : 2006-01-02
func TimeToStrWFormatDate(time time.Time) string {
	return time.Format("2006-01-02")
}
