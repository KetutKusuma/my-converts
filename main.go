package myconverts

import (
	"strconv"
	"time"
)

/// BORHH

// convert time to string with format : 2006-01-02 15:04:05
func TimeToStrWFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// convert time to string with format : 2006-01-02
func TimeToStrWFormatDate(time time.Time) string {
	return time.Format("2006-01-02")
}

// if you want to data with date today and time max
// ex : 2024-10-10 23:59:59
func TimeToMaxTime(tm time.Time) time.Time {
	return tm.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

// string to int without error if error return 0
func StrtoIntNoERR(st string) int {
	data, err := strconv.Atoi(st)
	if err != nil {
		return 0
	}
	return data
}
