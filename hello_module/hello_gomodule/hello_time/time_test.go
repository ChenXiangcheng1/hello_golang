package hellotime

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	// 时间常量
	time.Sleep(1 * time.Second)

	// 转换
	now := GetCurrentTime()
	fmt.Printf("%T %v\n", now, now)
	nowUTC := now.UTC()
	fmt.Printf("%T %v\n", nowUTC, nowUTC)
	nowStr := now.String()
	fmt.Printf("%T %v\n", nowStr, nowStr)
	fmtStr := "2006-01-02T15:04:05+0800"
	nowStr2 := now.Format(fmtStr) // 将now转换为字符串，Format()参数必须是"2006-01-02T15:04:05+0000 UTC" 12345
	fmt.Println(now, nowStr2)
	fmt.Printf("%T %v\n", nowStr2, nowStr2)
	nowS1 := now.Unix() // 秒
	fmt.Printf("%T %v\n", nowS1, nowS1)
	nowS2 := now.UnixNano() // 纳秒
	fmt.Printf("%T %v\n", nowS2, nowS2)
	nowS3 := Time2TimeStampMill(now)
	fmt.Printf("%T %v\n", nowS3, nowS3)

	timeStamp := int64(1581348090000)
	fmt.Println(TimeStamp2TimeSecond(timeStamp), TimeStamp2TimeMill(timeStamp), TimeStamp2TimeNano(timeStamp))

	// 创建
	myTime := time.Date(2020, 2, 10, 20, 1, 30, 0, TIME_LOCATION)
	fmt.Printf("%T %v\n", myTime, myTime)

	year, month, day := myTime.Date()
	fmt.Printf("%T %v\n", year, year)
	fmt.Printf("%T %v\n", month, month) // time.Month
	fmt.Printf("%T %v\n", day, day)
	fmt.Println(year, month, day)
	hour, minute, second := now.Clock()
	fmt.Println(hour, minute, second)
	fmt.Println(now.Weekday())

	// 时间比较计算(需要在同一个时区)
	fmt.Println(myTime.Before(now))
	fmt.Println(myTime.After(now))
	fmt.Println(myTime.Equal(now))

	// 时间计算
	timeDuration := now.Sub(myTime)
	fmt.Printf("%T %v\n", timeDuration, timeDuration)
	timeDuration_hours := timeDuration.Hours()
	fmt.Printf("%T %v\n", timeDuration_hours, timeDuration_hours)
	fmt.Println(timeDuration.Minutes(), timeDuration.Seconds())

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("%T %v\n", tomorrow, tomorrow)
	oneHourLater := tomorrow.Add(1 * time.Hour)
	fmt.Printf("%T %v\n", oneHourLater, oneHourLater)
}
