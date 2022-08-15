package gtools

import "fmt"

// Shell中输出彩色字体
var (
	//greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	//whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	//yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	//redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	//blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	//magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	//cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green = string([]byte{27, 91, 51, 50, 109})
	//white        = string([]byte{27, 91, 51, 55, 109})
	yellow = string([]byte{27, 91, 51, 51, 109})
	red    = string([]byte{27, 91, 51, 49, 109})
	blue   = string([]byte{27, 91, 51, 52, 109})
	//magenta      = string([]byte{27, 91, 51, 53, 109})
	//cyan         = string([]byte{27, 91, 51, 54, 109})
	reset = string([]byte{27, 91, 48, 109})
	//disableColor = false
)

func PrintGreen(str interface{}) {
	fmt.Print(green, str, reset)
}

func PrintlnGreen(str interface{}) {
	fmt.Println(green, str, reset)
}

func PrintfGreen(str interface{}) {
	fmt.Printf(green, str, reset)
}

func PrintYellow(str interface{}) {
	fmt.Print(yellow, str, reset)
}

func PrintlnYellow(str interface{}) {
	fmt.Println(yellow, str, reset)
}

func PrintfYellow(str interface{}) {
	fmt.Printf(yellow, str, reset)
}

func PrintRed(str interface{}) {
	fmt.Print(red, str, reset)
}

func PrintlnRed(str interface{}) {
	fmt.Println(red, str, reset)
}

func PrintfRed(str interface{}) {
	fmt.Printf(red, str, reset)
}

func PrintBlue(str interface{}) {
	fmt.Println(blue, str, reset)
}

func PrintlnBlue(str interface{}) {
	fmt.Println(blue, str, reset)
}

func PrintfBlue(str interface{}) {
	fmt.Printf(blue, str, reset)
}
