package gtools

import (
	"fmt"
	"time"
)

// GetDate 获取当前时间
func GetDate() string {
	now := time.Now()
	date := fmt.Sprint(now.Format("2006-01-02 15:04:05"))
	return date
}

