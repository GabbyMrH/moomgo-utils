/*
 * @PackageName: time
 * @Description: 常用time操作
 * @Author: Casso
 * @Date: 2022-01-28 10:20:33
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 17:06:07
 */

package time

import (
	"log"
	"math/rand"
	"time"
)

var (
	YYMMDDHHMMSS = "2006-01-02 15:04:05"
	YYMMDD       = "2006-01-02"
)

type DS struct {
	Ds string
	T  time.Time
}

// StringToTime 时间字符串转time
func StringToTime(timeStr string) time.Time {
	t, err := time.ParseInLocation(YYMMDDHHMMSS, timeStr, time.Local)
	if err != nil {
		// 时间解析，异常
		log.Fatalf("StringToTime ParseInLocation Err:%v", err)
		return t
	}

	return t
}

// TimeToStringFull time转字符串 2006-01-02 15:04:05
func TimeToStringFull(time time.Time) string {
	return time.Format(YYMMDDHHMMSS)
}

// TimeToStringYMD time转字符串 2006-01-02
func TimeToStringYMD(time time.Time) string {
	return time.Format(YYMMDD)
}

// GetTimestampForDay 获取当天 开始&结束 时间戳
func GetTimestampForDay() (beginTimeNum, endTimeNum int64) {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	beginTimeNum = t.Unix()
	endTimeNum = beginTimeNum + 86400
	return beginTimeNum, endTimeNum
}

// GetTimeUnix 获取秒级时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// GetTimeUnixMilli 获取毫秒级时间戳
func GetTimeUnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetNowTimeUnixBefore 获取传入时间到当前的时间差
func GetNowTimeUnixBefore(duration time.Duration) int64 {
	return time.Now().Add(-duration).Unix()
}

// GetTimeUnixAfter 获取传入时间 指定时间后的时间戳
func GetTimeUnixAfter(t time.Time, duration time.Duration) int64 {
	return t.Add(duration).Unix()
}

// GetNowTimeUnixAfter 获取当前时间加上传入时间的时间戳
func GetNowTimeUnixAfter(duration time.Duration) int64 {
	return time.Now().Add(duration).Unix()
}

// GetRandomRedisTimeOut 在指定时间上添加一个300秒的随机值
func GetRandomRedisTimeOut(base time.Duration) time.Duration {
	rand.Seed(time.Now().UnixNano())
	add := time.Duration(rand.Intn(5*60)) * time.Second
	return base + add
}

// GetEndDaySecond 获取当天还剩下多少秒
func GetEndDaySecond() int64 {
	t := time.Now()
	d := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return d.AddDate(0, 0, 1).Unix() - GetTimeUnix()
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []DS {
	var (
		d []DS
	)
	if len(YYMMDDHHMMSS) != len(sdate) {
		YYMMDDHHMMSS = YYMMDDHHMMSS[0:len(sdate)]
	}

	date, err := time.Parse(YYMMDDHHMMSS, sdate)
	if err != nil {
		// 时间解析，异常
		return nil
	}

	date2, err := time.Parse(YYMMDDHHMMSS, edate)
	if err != nil {
		// 时间解析，异常
		return nil
	}

	if date2.Equal(date) {
		d = append(d, DS{
			Ds: date.Format(YYMMDDHHMMSS),
			T:  date2,
		})
		return d
	}

	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return nil
	}

	// 输出日期格式固定
	date2Str := date2.Format(YYMMDDHHMMSS)
	d = append(d, DS{
		Ds: date.Format(YYMMDDHHMMSS),
		T:  date2,
	})

	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(YYMMDDHHMMSS)
		d = append(d, DS{
			Ds: date.Format(YYMMDDHHMMSS),
			T:  date,
		})
		if dateStr == date2Str {
			break
		}
	}

	return d
}

// GetTimesDiffer 获取相差时间秒数
func GetTimesDiffer(stime, etime string) (res int64, errs error) {
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", stime, time.Local)
	t2, err2 := time.ParseInLocation("2006-01-02 15:04:05", etime, time.Local)
	if err == nil && err2 == nil && t1.Before(t2) {
		res = t2.Unix() - t1.Unix()
		errs = nil
	} else {
		res = 0
		errs = err
	}
	return
}
