package tools

import "fmt"

// 实现输出带颜色字体功能：fmt.Print、fmt.Println

var (
	red    = "\033[31m"
	green  = "\033[32m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

// 颜色输出fmt.Print

type PrintColor struct {
	text interface{}
}

func NewPrint(text interface{}) PrintColor {
	var p PrintColor
	p.text = text
	return p
}

func (p PrintColor) Red() {
	fmt.Print(red, p.text, reset)
}

func (p PrintColor) Green() {
	fmt.Print(green, p.text, reset)
}

func (p PrintColor) Blue() {
	fmt.Print(blue, p.text, reset)
}

func (p PrintColor) Yellow() {
	fmt.Print(yellow, p.text, reset)
}

// 颜色输出fmt.Println，前面会多出一个空格，可以使用fmt.Print("xxx\n")形式处理空格问题

type PrintlnColor struct {
	text interface{}
}

func NewPrintln(text interface{}) PrintlnColor {
	var p PrintlnColor
	p.text = text
	return p
}

func (p PrintlnColor) Red() {
	fmt.Println(red, p.text, reset)
}

func (p PrintlnColor) Green() {
	fmt.Println(green, p.text, reset)
}

func (p PrintlnColor) Blue() {
	fmt.Println(blue, p.text, reset)
}

func (p PrintlnColor) Yellow() {
	fmt.Println(yellow, p.text, reset)
}
