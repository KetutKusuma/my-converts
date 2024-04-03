package main

import "time"

func StrToTimeWFormat(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func TimeToStrWFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
