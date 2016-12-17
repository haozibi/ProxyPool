package proxy

import (
	"flag"
	gg "github.com/haozibi/gglog"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var proxyAllList chan string = make(chan string, 300)
var proxyOkList chan string = make(chan string, 100)
var proxyNum int = 0
var ProxyProNum int = 50
var mu, mu2 sync.Mutex
var rmu, rmu2 sync.RWMutex
var TestUrl = "http://www.baidu.com"

func init() {
	flag.Parse()
	defer gg.Flush()

	// 设置log输出路径
	//gg.SetLogDir("log")

	// 设置控制台输出级别，比此级别大的都会在控制台输出
	// DEBUG < INFO < WARING < ERROR < FATAL , 默认ERROR级别
	gg.SetOutLevel("INFO")
	// 设置 console 输出是否精简，默认完整输出
	gg.SetOutSimple(true)

	// 一开始就进行代理获取
	checkProxy()
}

func GetProxy() (proxyUri string) {
	rmu2.RLock()
	uri := <-proxyOkList
	rmu2.RUnlock()
	return uri
}

func Setting(flag bool) {
	if flag == true {
		// 设置控制台输出级别，比此级别大的都会在控制台输出
		// DEBUG < INFO < WARING < ERROR < FATAL , 默认ERROR级别
		gg.SetOutLevel("DEBUG")
		//获取代理过程输出
		go func() {
			for {
				gg.Debugf("proxyOkList: %d", len(proxyOkList))
				gg.Debugf("proxyAllList: %d", len(proxyAllList))
				gg.Debugf("All: %d", proxyNum)
				time.Sleep(2 * time.Second)
			}
		}()
	}
}

func getProxyList() []string {
	var tmpList = make([]string, 300)
	for _, v := range proxyOne() {
		tmpList = append(tmpList, v)
	}
	for _, v := range proxyTwo() {
		tmpList = append(tmpList, v)
	}
	for _, v := range proxyThree() {
		tmpList = append(tmpList, v)
	}
	return tmpList
}

func checkProxy() {
	gg.Infoln("Start")
	go func() {
		for {
			if len(proxyAllList) < minProxyLen {
				mu.Lock()
				for _, v := range getProxyList() {
					if v != "" {
						gg.Debugln(v)
						proxyAllList <- v
					}
				}
				mu.Unlock()
				gg.Infoln("Proxy Add URL Over")
				for i := 0; i < ProxyProNum; i++ {
					go doCheckProxy()
				}
			}
			// 60秒检测是否需要添加新的未筛选代理IP
			time.Sleep(60 * time.Second)
		}
	}()
}

func doCheckProxy() {
	for {
		rmu.RLock()
		uri := <-proxyAllList
		rmu.RUnlock()
		gg.Debugf("uri:%v", uri)
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + uri) //根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
		}
		transport := &http.Transport{Proxy: proxy}
		client := &http.Client{Transport: transport}
		// client := &http.Client{Transport: transport, Timeout: 10 * time.Second}
		req, _ := http.NewRequest(
			"GET",
			TestUrl,
			nil,
		)
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("Accept-Encoding", "text/html")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		resp, err := client.Do(req)
		if err == nil {
			if resp.StatusCode == 200 {
				mu2.Lock()
				proxyOkList <- uri
				proxyNum = proxyNum + 1
				gg.Debugln("ok", uri, resp.Request.URL)
				mu2.Unlock()
			}
		} else {
			gg.Debugln("err", uri)
		}
	}
}
