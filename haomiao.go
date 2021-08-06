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

// InputInt 获取用户输入数字
func InputInt(info string) (num int) {
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

// Shell中输出彩色字体
var (
	//greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	//whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	//yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	//redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	//blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	//magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	//cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green        = string([]byte{27, 91, 51, 50, 109})
	//white        = string([]byte{27, 91, 51, 55, 109})
	yellow       = string([]byte{27, 91, 51, 51, 109})
	red          = string([]byte{27, 91, 51, 49, 109})
	blue         = string([]byte{27, 91, 51, 52, 109})
	//magenta      = string([]byte{27, 91, 51, 53, 109})
	//cyan         = string([]byte{27, 91, 51, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	//disableColor = false
)

func PrintlnGreen(str string){
	fmt.Println(green, str, reset)
}

func PrintlnYellow(str string){
	fmt.Println(yellow, str, reset)
}

func PrintlnRed(str string){
	fmt.Println(red, str, reset)
}

func PrintlnBlue(str string){
	fmt.Println(blue, str, reset)
}