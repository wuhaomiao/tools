package tools

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// CMD shell实时输出, ScanLines以行输出。Words以单个词输出
func CMD(shell string) error {
	cmd := exec.Command("/bin/bash", "-c", shell)
	reader, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println("shell执行出错：", err)
		return err
	}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("cmd.Wait()执行出错：", err)
		return err
	}
	return nil
}

// NewCMD shell实时输出
func NewCMD(shell string) error {
	cmd := exec.Command("/bin/bash", "-c", shell)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("shell命令执行失败: ", err)
		return err
	}
	return nil
}

// IsDirExists 判断目录是否存在, 不存刚创建目录
func IsDirExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("无法创建目录： %v", err)
		}
	}
	return nil
}
