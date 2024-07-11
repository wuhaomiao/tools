package tools

import (
	"fmt"
	"time"
)

// GetDate 获取当前日期
func GetDate() string {
	now := time.Now()
	date := fmt.Sprint(now.Format("2006-01-02"))
	return date
}

// GetDateString 获取当前日期
func GetDateString() string {
	now := time.Now()
	date := fmt.Sprint(now.Format("20060102"))
	return date
}

// GetTime 获取当前时间
func GetTime() string {
	now := time.Now()
	date := fmt.Sprint(now.Format("2006-01-02 15:04:05"))
	return date
}
