package proxy

import (
	"errors"
	"flag"
	"regexp"
	"sync"
	"time"

	gg "github.com/haozibi/gglog"
)

var (
	testUrl      = "https://www.baidu.com"
	regexIPPort  = `^(25[0-5]|2[0-4]\d|1\d\d|\d\d|\d)\.(25[0-5]|2[0-4]\d|1\d\d|\d\d|\d)\.(25[0-5]|2[0-4]\d|1\d\d|\d\d|\d)\.(25[0-5]|2[0-4]\d|1\d\d|\d\d|\d)(:(\d\d\d\d|\d\d\d|\d\d|\d))$`
	regexTestUrl = `(ht|f)tp(s?)\://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`
	mutex        sync.Mutex
	rmutex       sync.RWMutex
	maxListLen   = 100
)
var Proxys []*proxy

var proxyOKList = make(chan string, maxListLen)

func addProxy(p *proxy) {
	if !p.isAvailable {
		gg.Infof("Proxy [%v] not available\n", p.name)
		return
	}
	Proxys = append(Proxys, p)
}

type proxy struct {
	name         string
	isAvailable  bool
	allList      []string
	timeInterval int                      // 每次获取的时间间隔，以秒为单位
	getList      func() ([]string, error) // 返回可用的 ip 列表 123.123.123.123:1024
}

func (p *proxy) init() {
	p.allList = make([]string, 0)
}

// 检查代理是否可用
// todo: 增加多个测试链接
func (p *proxy) checkProxy() {
	// fmt.Println(runtime.NumGoroutine())
	if len(p.allList) == 0 {
		gg.Infof("Proxy [%v] not get ip:port", p.name)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(p.allList))
	for _, v := range p.allList {
		go func(uri string, w *sync.WaitGroup) {
			defer func() {
				w.Done()
			}()

			if m, _ := regexp.MatchString(regexIPPort, uri); !m {
				return
			}
			_, ok, _ := HttpFunc(testUrl, uri, "GET")
			if ok {
				addList(uri)
			}
			gg.Debugf("Proxy [%v] %v => %v\n", p.name, uri, ok)
			return
		}(v, &wg)
	}
	wg.Wait()
}

func addList(s string) {
	mutex.Lock()
	proxyOKList <- s
	mutex.Unlock()
}

func getList() string {
	rmutex.Lock()
	s := <-proxyOKList
	rmutex.Unlock()
	return s
}

func init() {
	flag.Parse()
	defer gg.Flush()

	gg.SetOutLevel("INFO")
	gg.SetOutType("SIMPLE")
	gg.SetPrefix("[ProxyPool] ")
}

func StartProxy() {
	go startProxy()
}

func startProxy() {
	if len(Proxys) == 0 {
		gg.Errorf("Not found available proxy server\n")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(Proxys))
	for _, v := range Proxys {
		if !v.isAvailable {
			continue
		}
		gg.Infof("Add proxy [%v] success\n", v.name)
		go func(p *proxy, w *sync.WaitGroup) {
			defer w.Done()
			for {
				p.init()
				s, err := p.getList()
				if err != nil {
					gg.Infof("Proxy [%v] waiting %v second, due to error,%v\n", p.name, p.timeInterval, err)
					time.Sleep(time.Duration(p.timeInterval) * time.Second)
					continue
				}
				p.allList = append(p.allList, s...)
				p.checkProxy()
				gg.Infof("Proxy [%v] waiting %v second...\n", p.name, p.timeInterval)
				time.Sleep(time.Duration(p.timeInterval) * time.Second)
			}
		}(v, &wg)
	}
	wg.Wait()
}

func SetTestUrl(uri string) error {
	if m, _ := regexp.MatchString(regexTestUrl, uri); m {
		testUrl = uri
		gg.Infoln("Set test url success")
		return nil
	}
	return errors.New("Test url shoud error")
}

// 获取一个可用的代理
func GetProxy() string {
	return getList()
}
