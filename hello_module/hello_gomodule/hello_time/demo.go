package hello_time

import (
	"errors"
	"fmt"
	"time"
)

/*
对time包的功能封装
*/

var TIME_LOCATION *time.Location

// 初始化时区
func init() {
	var err error
	TIME_LOCATION, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("%v %T\n", err, err)
		panic(err)
	}
}

func GetCurrentTime() time.Time {
	now := time.Now()
	fmt.Printf("demo.go: %T %v\n", now, now) // time.Time
	ret := now.In(TIME_LOCATION)             // 转换时区
	fmt.Printf("demo.go: %T %v\n", ret, ret)
	return ret
}

// 判断两个日期是否是同一天
func DateEqual(day1, day2 time.Time) bool {
	day1 = day1.In(TIME_LOCATION)
	day2 = day2.In(TIME_LOCATION)
	y1, m1, d1 := day1.Date()
	y2, m2, d2 := day2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// 获取指定时间的毫秒级别的时间戳, now.Unix()获取秒级别, now.UnixNano()获取纳秒级别
func Time2TimeStampMill(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 秒时间戳转时间
func TimeStampSecond2Time(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0)
}

// 毫秒时间戳转时间
func TimeStampMill2Time(timeStamp int64) time.Time {
	return time.Unix(0, timeStamp*1e6)
}

// 纳秒时间戳转时间
func TimeStampNano2Time(timeStamp int64) time.Time {
	return time.Unix(0, timeStamp)
}

const (
	MyNano      = "2006-01-02 15:04:05.000000000"
	MyMicro     = "2006-01-02 15:04:05.000000"
	MyMil       = "2006-01-02 15:04:05.000"
	MySec       = "2006-01-02 15:04:05"
	MyCST       = "2006-01-02 15:04:05 +0800 CST"
	MyUTC       = "2006-01-02 15:04:05 +0000 UTC"
	MyDate      = "2006-01-02"
	MyTime      = "15:04:05"
	FBTIME      = "2006-01-02T15:04:05+0800" // Facebook
	APPTIME     = "2006-01-02T15:04:05.000"  // Application
	TWITTERTIME = "2006-01-02T15:04:05Z"     // Twitter
)

func getTimeDefault() time.Time {
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", "", TIME_LOCATION)
	// return time.Time{}
	return tm
}

// 字符串转time.Time
func TimeStr2Time(timeStr string) (time.Time, error) {
	fmtStrArray := []string{
		MyDate, MyTime,
		time.RFC3339,
		time.RFC3339Nano,
	}
	var tm time.Time
	var err error
	for _, fmtStr := range fmtStrArray {
		tm, err = time.ParseInLocation(fmtStr, timeStr, TIME_LOCATION)
		if err != nil {
			continue
		} else {
			break
		}
	}
	if tm == getTimeDefault() {
		return tm, errors.New("时间字符串格式错误")
	}
	return tm, nil
}
