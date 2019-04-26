package goutil

import (
	"time"
)

const (
	DATE_FORMAT = "2006-01-02"
	DATEHM_FORMAT = "2006-01-02 15:04"
	DATEHMI_FORMAT = "2006-01-02 15:04:05"
)
//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day() + 1)
	return GetZeroTime(d)
}
//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

//日期转时间
func StringToTime(toBeCharge string) time.Time {
	timeLayout := DATEHMI_FORMAT
	//loc,_ := time.LoadLocation("local")

	theTime,_ := time.ParseInLocation(timeLayout,toBeCharge,time.Local)

	return theTime
}

//根据指定格式对时间戳转换
func GetDateFormat(timestamp int64, format string) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

//时间戳转日期，没有时分秒
func GetDate(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(DATE_FORMAT)
}

//时间戳转日期
func GetDateMHI(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(DATEHMI_FORMAT)
}

//时间戳转日期，没有秒
func GetDateMH(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(DATEHM_FORMAT)
}

//日期转时间戳,包含时分秒，参数格式必须相同，如2006-01-02 15:04:05
func GetTimeParse(times string) int64 {
	if "" == times {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation(DATEHMI_FORMAT, times, loc)
	return parse.Unix()
}

//根据给的日期和格式转时间戳，参数格式如2006-01-02 15:04:05
func GetDateParse(date string,tpl string) int64 {
	if "" == date {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation(tpl, date, loc)
	return parse.Unix()
}

//根据给的日期和格式转换
func DateTodate(date string,tpl string,tpl2 string) string {
	if "" == date {
		return ""
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation(tpl, date, loc)
	return parse.Format(tpl2)
}
