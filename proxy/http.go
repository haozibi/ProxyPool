package proxy

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func HttpFunc(uri, proxyUri, method string) ([]byte, bool, error) {
	client := &http.Client{}
	if len(proxyUri) != 0 {
		proxyFunc := func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + proxyUri) //根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
		}
		transport := &http.Transport{Proxy: proxyFunc}
		client = &http.Client{Transport: transport}
	}

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return []byte(""), false, err
	}

	req.Header.Set("User-Agent", getRandomUserAgent())

	resp, err := client.Do(req)
	// 请求失败
	if err != nil {
		return []byte(""), false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte(""), false, fmt.Errorf("%v status code not equal 200", uri)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), false, err
	}

	return body, true, nil
}

func getRandomUserAgent() string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgents[r.Intn(len(userAgents))]
}

var userAgents = [...]string{
	"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
}
