package util

import (
	"fmt"
	"net"
	"strings"
)

// IsPortAvailable 检查指定端口是否可用
func IsPortAvailable(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return false // 端口被占用或无法监听
	}
	_ = ln.Close() // 关闭监听，释放端口
	return true    // 端口可用
}

// IsEmpty 判断字符串是否为空或全是空格
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotEmpty 判断字符串是否不为空且不全是空格
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// HasEmpty 判断多个字符串中是否存在空字符串
func HasEmpty(strs ...string) bool {
	for _, s := range strs {
		if IsNotEmpty(s) {
			return false
		}
	}
	return true
}

// HasNotEmpty 判断多个字符串中是否都不为空
func HasNotEmpty(strs ...string) bool {
	for _, s := range strs {
		if IsEmpty(s) {
			return false
		}
	}
	return true
}
