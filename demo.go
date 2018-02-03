package main

import (
	"flag"

	"github.com/haozibi/ProxyPool/proxy"
	gg "github.com/haozibi/gglog"
)

func main() {
	flag.Parse()
	defer gg.Flush()

	gg.SetOutLevel("INFO")
	// gg.SetOutType("SIMPLE")
	gg.SetPrefix("[ProxyPool] ")

	// 设置代理测试网站
	err := proxy.SetTestUrl("http://www.weibo.com")
	if err != nil {
		panic(err)
	}
	// 开始测试代理
	// proxy.StartProxy()
	// for {
	// 	gg.Infoln("get ==>", proxy.GetProxy())
	// }
	proxy.StartProxyByWeb()
}
