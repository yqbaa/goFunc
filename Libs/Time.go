package Libs

import (
	"time"
	"strings"
)

type Time struct {
}

/*
 * 返回时间戳
 */
func (this *Time)Time() int64{
	return time.Now().Unix()
}

/*
 * 返回微秒时间戳
 */
func (this *Time)Microtime() int64{
	return time.Now().UnixNano()
}

/*
*格式化时间
 */
func (this *Time)Date(format string,t int64)string{
	format = this.formatReal(format)
	return time.Unix(t, 0).Format(format)
}

func (this *Time)formatReal(format string) string{
	var Format map[string]string
	Format = make(map[string]string)
	Format["Y"] = "2006"
	Format["y"] = "06"
	Format["m"] = "01"
	Format["d"] = "02"
	Format["H"] = "15"
	Format["i"] = "04"
	Format["s"] = "05"
	for k,v := range Format{
		format = strings.Replace(format , k, v, -1)
	}
	return format
}

/*
 * 字符串转成时间戳
 * 比如2018-09-09  2018-09-09 09:19 etc
 */
func (this *Time)Strtotime(str string)int64{
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, str, loc) //使用模板在对应时区转化为time.time类型
	return theTime.Unix()                                            //转化为时间戳 类型是int64
}