package main

import (
	"github.com/haozibi/ProxyPool/proxy"
	gg "github.com/haozibi/gglog"
)

func main() {
	// 设置代理测试网站
	err := proxy.SetTestUrl("http://www.weibo.com")
	if err != nil {
		panic(err)
	}
	// 开始测试代理
	proxy.StartProxy()
	for {
		gg.Infoln("get ==>", proxy.GetProxy())
	}
}
