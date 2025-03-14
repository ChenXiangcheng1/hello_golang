package hellotime

import (
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
	fmt.Printf("%T %v\n", now, now) // time.Time
	ret := now.In(TIME_LOCATION)    // 转换时区
	fmt.Printf("%T %v\n", ret, ret)
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

// 获取指定时间的毫秒级别的时间戳
func Time2TimeStampMill(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 时间戳转时间
func TimeStamp2TimeSecond(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0)
}

func TimeStamp2TimeMill(timeStamp int64) time.Time {
	return time.Unix(0, timeStamp*1e6)
}

func TimeStamp2TimeNano(timeStamp int64) time.Time {
	return time.Unix(0, timeStamp)
}
