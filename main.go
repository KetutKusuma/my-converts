package myconverts

import "time"

/// BORHH

// convert time to string with format : 2006-01-02 15:04:05
func TimeToStrWFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// convert time to string with format : 2006-01-02
func TimeToStrWFormatDate(time time.Time) string {
	return time.Format("2006-01-02")
}

// convert string to time with format : 2006-01-02 15:04:05
func StrToTimeWFormat(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

// convert string to time with format : 2006-01-02
func StrToTimeWFormatDate(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

// convert to str from date with full format : 2006-01-02T15:04:05+08:00
func StrToTimeFullFormat(str string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05+08:00", str)
}

// convert string to time with max hour, minute and second
// needs by filter date ex : "2024-11-01"
// the output : 2024-11-01 23:59:59
func StrToTimeWFrmtDateMaxTime(str string) (time.Time, error) {
	tim, err := time.Parse("2006-01-02", str)
	if err != nil {
		return time.Time{}, err
	}
	timeMax := tim.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	return timeMax, nil
}

// if you want to data with date today and time max
// ex : 2024-10-10 23:59:59
func TimeToMaxTime(tm time.Time) time.Time {
	return tm.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}
