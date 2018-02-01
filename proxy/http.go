package proxy

import (
	"io/ioutil"
	"net/http"
	"net/url"
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

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")

	resp, err := client.Do(req)
	// 请求失败
	if err != nil {
		return []byte(""), false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), false, err
	}

	if resp.StatusCode == http.StatusOK {
		return body, true, nil
	}

	return body, false, nil
}
