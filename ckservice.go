package tools

import (
	"fmt"
	"os/exec"
	"strings"
)

// CKService 检查是否安装某个服务，返回布尔值
func CKService(serviceName string) bool {
	cmd := fmt.Sprintf(`systemctl status %s|grep Active| awk '{print $2}'`, serviceName)
	if status, _ := exec.Command("/bin/bash", "-c", cmd).Output(); strings.TrimSpace(string(status)) != "active" {
		return false
	}
	return true
}
