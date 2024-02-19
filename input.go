package tools

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
			Print("错误: 你输入的不是数字\n").Yellow()
			continue // 如果输入的不是数字, 返回-1
		}
		if num < 0 {
			Print("错误: 请输入正整数\n")
			continue
		}
		return num
	}
	return -1
}

// InputIntMax 用户输入范围内正整数。 -1: 代表用户输入有误。-2: 代表用户直接回车。
func InputIntMax(userInput string, minInt, maxInt int) int {
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
			Print("错误: 你输入的不是数字\n").Yellow()
			continue // 如果输入的不是数字, 返回-1
		}
		if num < minInt {
			Print("错误: 不能小于最小值\n").Red()
			continue
		}
		if num > maxInt {
			Print("错误: 不能大于最大值\n").Red()
		}
		return num
	}
	return -1
}

// Choose 用户输入[y/n], 返回bool值，true为yes，false为no.
func Choose(userInput string) bool {
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
