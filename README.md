# proxy

[![Build Status](https://travis-ci.org/haozibi/ProxyPool.svg?branch=master)](https://travis-ci.org/haozibi/ProxyPool) ![](https://img.shields.io/badge/language-go-blue.svg)

**代理 IP 筛选**，由于网络上获取的代理不尽人意，所以写了个程序对代理进行筛选。

个人感觉代理的时效性，所以没有进行持久化设计。*分布式进行持久化设计比较好*


## 安装

> go get -u -v github.com/haozibi/ProxyPool

**注意：**可以参考 `proxy/web.66ip.cn.go` 和 `proxy/web.test.go` 编写专属获取代码获取更多待筛选的代理 IP

## 示例

```go
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
```
## 注意

代理 IP 测试暂时只支持 `http://`，不支持 `https://`

## 免费IP参考

**只是测试网站，请不要滥用！**

* 66ip.cn [http://m.66ip.cn/index.html](http://m.66ip.cn/index.html)

## 截图
![](https://i.loli.net/2018/02/01/5a732dd9caebe.jpg)
