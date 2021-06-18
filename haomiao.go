package haomiao

import (
	"fmt"
	"time"
)

// GetDate 获取系统时间，返回格式化后的日期
func GetDate() string {
	now := time.Now()
	nowTime := fmt.Sprintln(now.Format("2006-01-02 15:04:05"))
	return nowTime
}

// Line 打印分隔线
func Line() {
	for i := 0; i < 6; i++ {
		fmt.Print("==========")
	}
	fmt.Println("")
}
