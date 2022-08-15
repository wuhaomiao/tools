package gtools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// InputStr 用户输入字符串
func InputStr(userInput string) string {
	fmt.Print(userInput)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("输入有误: ", err)
	}
	text = strings.TrimSpace(text)
	return text
}

// InputInt 用户输入正整数。 -1: 代表用户输入有误。-2: 代表用户直接回车。
func InputInt(userInput string) int {
	for i := 0; i < 5; i++ {
		fmt.Print(userInput)
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("输入有误: ", err)
		}
		text = strings.TrimSpace(text)
		if text == "" {
			return -2 // 如果直接回车，返回-2
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			PrintlnYellow("错误: 你输入的不是数字")
			continue // 如果输入的不是数字, 返回-1
		}
		if num < 0 {
			PrintlnRed("错误: 请输入正整数")
			continue
		}
		return num
	}
	return -1
}

// ChooseYes 用户输入[y/n], 返回bool值，true为yes，false为no.
func ChooseYes(userInput string) bool {
	for i := 0; i < 5; i++ {
		fmt.Print(userInput)
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("输入有误: ", err)
		}
		text = strings.TrimSpace(text)
		choose := strings.ToUpper(text)
		if choose == "Y" || choose == ("YES") {
			return true
		} else if choose == "N" || choose == "NO" {
			return false
		} else {
			continue
		}
	}
	return false
}
