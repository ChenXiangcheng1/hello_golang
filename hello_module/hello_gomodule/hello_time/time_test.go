package hello_time

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	getTimeDefault()
}

// 转换
func TestConvert(t *testing.T) {
	now := GetCurrentTime()
	fmt.Printf("Now: %T %v\n", now, now) // China Standard Time
	nowUTC := now.UTC()
	fmt.Printf("UTC: %T %v\n", nowUTC, nowUTC) // Coordinated Universal Time, monotonic clock单调时钟

	// time.Time转为String
	nowStr := now.String()
	fmt.Printf("str: %T %v\n", nowStr, nowStr)
	fmtStr := "2006-01-02T15:04:05+0800" // +0800是字符串
	nowStr2 := now.Format(fmtStr)        // 将now转换为字符串，Format()参数必须是"2006-01-02T15:04:05+0000 UTC" 12345
	fmt.Printf("str2: %T %v\n", nowStr2, nowStr2)

	// String转time.Time
	timeStr := "2020-02-10T20:01:30+0800"
	fmtStr = "2006-01-02T15:04:05+0800" // +0800是字符串
	// t1, err := time.Parse(fmtStr, timeStr) // 注意t和time都是已有变量
	t1, err := time.ParseInLocation(fmtStr, timeStr, TIME_LOCATION)
	if err != nil {
		fmt.Printf("%T %v\n", err, err) // *time.ParseError
		t.Errorf("time.Parse error: %v", err)
	} else {
		fmt.Printf("Time: %T %v\n", t1, t1)
	}

	timeStr = "2020-02-10T20:01:30+0800"
	t2, err2 := TimeStr2Time(timeStr)
	if err2 != nil {
		fmt.Printf("%T %v\n", err2, err2)
	} else {
		fmt.Printf("Time: %T %v\n", t2, t2)
	}

	// time.Time转时间戳
	secd := now.Unix() // 秒时间戳
	fmt.Printf("Second: %T %v\n", secd, secd)
	nano := now.UnixNano() // 纳秒时间戳
	fmt.Printf("Nano: %T %v\n", nano, nano)
	mill := Time2TimeStampMill(now) // 毫秒时间戳
	fmt.Printf("Mill: %T %v\n", mill, mill)

	// 时间戳转time.Time
	t1 = TimeStampSecond2Time(secd)
	fmt.Printf("Time: %T %v\n", t1, t1)
	t1 = TimeStampNano2Time(nano)
	fmt.Printf("Time: %T %v\n", t1, t1)
	t1 = TimeStampMill2Time(mill)
	fmt.Printf("Time: %T %v\n", t1, t1)
}

// time.Time的计算和比较
func TestCalc(t *testing.T) {
	tm := GetCurrentTime()
	tomorrow := tm.AddDate(0, 0, 1)
	tomorrow = tomorrow.Add(1 * time.Hour)
	timeDuration := tomorrow.Sub(tm)
	fmt.Printf("%T %v\n", timeDuration, timeDuration) // time.Duration

	timeDuration_hours := timeDuration.Hours()
	fmt.Printf("%T %v\n", timeDuration_hours, timeDuration_hours)
	timeDuration_minutes := timeDuration.Minutes()
	fmt.Printf("%T %v\n", timeDuration_minutes, timeDuration_minutes)

	fmt.Println(tm.Before(tomorrow))
	fmt.Println(tm.Equal(tomorrow))
	fmt.Println(tm.After(tomorrow))
}

// 时间常量
func TestConst(t *testing.T) {
	v := time.Nanosecond
	fmt.Printf("%T %v\n", v, v) // time.Duration int64(1)
}

// 定义
func TestDefine(t *testing.T) {
	// 定义
	mt := time.Date(2025, 3, 1, 16, 1, 07, 30, TIME_LOCATION)
	fmt.Printf("%T %v\n", mt, mt)
	// ts := int64(1581348090000)  // 时间戳

	// 属性和方法
	year, month, day := mt.Date()
	fmt.Printf("%T %v\n", year, year)
	fmt.Printf("%T %v\n", month, month) // time.Month
	fmt.Printf("%T %v\n", day, day)

	hour, minute, second := mt.Clock()
	fmt.Printf("%T %v\n", hour, hour)
	fmt.Printf("%T %v\n", minute, minute)
	fmt.Printf("%T %v\n", second, second)

	hour = mt.Hour()
	fmt.Printf("%T %v\n", hour, hour)

	weekday := mt.Weekday()
	fmt.Printf("%T %v\n", weekday, weekday)
}

func TestMain(t *testing.T) {
	time.Sleep(1 * time.Second)

	mt := GetCurrentTime()
	fmt.Printf("%T %v\n", mt, mt)
}
