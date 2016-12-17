# proxy

**免费代理ip筛选**

由于网络上免费获取的代理不尽人意，所以写了个程序对免费代理进行筛选，暂时支持三个免费代理网站

个人感觉代理的时效性，所以没有进行持久化设计。*分布式进行持久化设计比较好*

*多线程效率非常高*

> **空的 for 循环会导致 CPU 剧增**

## 示例
```
package main

import (
	"fmt"
	proxy "coding.net/haozibi/ProxyPool"
)

func main() {
	// 是否输出调试信息，true输出，false不输出。默认不输出
	proxy.Setting(true)

	// 自定义ip测试连接，默认 http://www.baidu.com
	proxy.TestUrl = "http://www.baidu.com"

	// 自定义线程数，默认50
	proxy.ProxyProNum = 100

	for {
		// GetProxy() 返回 175.171.246.195:8118 格式
		fmt.Println("Get ", proxy.GetProxy())
	}
}

```
## 注意

当可用 IP 达到 100 个则停止筛选，一旦少于 100 即立刻继续筛选

## 免费IP参考

共使用了3个免费网站

* 西刺代理 [http://api.xicidaili.com/free2016.txt](http://api.xicidaili.com/free2016.txt)
* 代理66 [http://www.66ip.cn/](http://www.66ip.cn/)
* 快代理IP [http://www.kuaidaili.com/](http://www.kuaidaili.com/)

## 截图
![](https://ooo.0o0.ooo/2016/12/14/58514ffff140b.png)
