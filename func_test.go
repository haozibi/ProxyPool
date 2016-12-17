package proxy

import (
	"fmt"
	"testing"
)

func TestGetProxy(t *testing.T) {
	// 是否输出调试信息，true输出，false不输出。默认不输出
	Setting(true)

	// 自定义ip测试连接，默认 http://www.baidu.com
	TestUrl = "http://www.baidu.com"

	// 自定义线程数，默认50
	ProxyProNum = 100

	for {
		// GetProxy() 返回 175.171.246.195:8118 格式
		fmt.Println("Get ", GetProxy())
	}
}
