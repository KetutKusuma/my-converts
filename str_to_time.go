package myconverts

import (
	"errors"
	"time"
)

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

// convert to str from date with full format : 2006-01-02 15:04:05 08:00
func StrToTimeFullFormatNoWordInside(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 08:00", str)
}

// convert to str from date with full format : 2006-01-02T15:04:05 08:00
func StrToTimeFullFormatNoPlus(str string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05 08:00", str)
}

// needs by filter date ex : "2009-11-10T23:00:00Z"
// the output : 2009-11-10 23:00:00
func StrToTimeWFrmtDateWithTZ(str string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z"
	tim, err := time.Parse(layout, str)
	return tim, err
}

// can be check all format of string like :
// 2006-01-02T15:04:05 08:00 //
// 2006-01-02T15:04:05+08:00 //
// 2006-01-02 //
// 2006-01-02 15:04:05 //
// 2009-11-10T23:00:00Z
func StrToTimeWAllFormatCheck(str string) (*time.Time, error) {
	timee, err := StrToTimeWFormat(str)
	if err == nil {

		return &timee, nil
	}
	// fmt.Println("str date StrToTimeWFormat failed, continue to StrToTimeWFormatDate")
	timee, err = StrToTimeWFormatDate(str)
	if err == nil {
		return &timee, nil
	}

	// fmt.Println("str date StrToTimeWFormatDate failed, continue to StrToTimeFullFormat")
	timee, err = StrToTimeFullFormat(str)
	if err == nil {
		return &timee, nil
	}

	// fmt.Println("str date StrToTimeFullFormat failed, continue to StrToTimeFullFormatNoPlus")
	timee, err = StrToTimeFullFormatNoPlus(str)
	if err == nil {
		return &timee, nil
	}

	// fmt.Println("str date StrToTimeFullFormatNoPlus failed, continue to StrToTimeFullFormatNoWordInside")
	timee, err = StrToTimeFullFormatNoWordInside(str)
	if err == nil {
		return &timee, nil
	}

	return nil, errors.New("can not convert str to time")
}
