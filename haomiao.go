package haomiao

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetDate 获取系统时间，返回格式化后的日期
func GetDate() string {
	now := time.Now()
	nowTime := fmt.Sprintln(now.Format("2006-01-02 15:04:05"))
	return nowTime
}

// InputStr 获取用户输入字符串
func InputStr(info string) (text string) {
	fmt.Print(info)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取有误: ", err)
		return
	}
	text = strings.TrimSpace(text)
	return text
}

// 获取用户输入数字
func inputInt(info string) (num int) {
	fmt.Print(info)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取有误: ", err)
		return
	}
	text = strings.TrimSpace(text)
	num, err = strconv.Atoi(text)
	if err != nil {
		fmt.Println("转换为Int类型出错：", err)
	}
	return num
}

// Line 打印分隔线
func Line() {
	for i := 0; i < 6; i++ {
		fmt.Print("==========")
	}
	fmt.Println("")
}
