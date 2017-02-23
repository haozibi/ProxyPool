package proxy

import (
	//"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	gg "github.com/haozibi/gglog"
)

// 每个网站一次获取100个，共300个

func proxyOne() []string {
	resp, err := http.Get(proxyOneUri)
	if err != nil {
		gg.Errorln("Error: ", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return strings.Split(string(body), "\r\n")
}

func proxyTwo() []string {
	resp, err := http.Get(proxyTwoUri)
	if err != nil {
		gg.Errorln("Error: ", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	reg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}`)
	return reg.FindAllString(string(body), -1)
}

func proxyThree() []string {
	var tmpList = make([]string, 100)
	for i := 1; i <= 10; i++ {
		client := &http.Client{}
		req, _ := http.NewRequest(
			"GET",
			proxyThreeUri+strconv.Itoa(i),
			nil,
		)
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("Accept-Encoding", "text/html")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		resp, err := client.Do(req)
		if err != nil {
			gg.Errorln("Error: ", err)
			return nil
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		regIP := regexp.MustCompile(`<td data-title="IP">\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}</td>`)
		regPort := regexp.MustCompile(`<td data-title="PORT">\d{1,5}</td>`)
		allIP := regIP.FindAllString(string(body), -1)
		allPort := regPort.FindAllString(string(body), -1)
		for i := 0; i < 10; i++ {
			str := fmt.Sprintf("%v:%v", allIP[i][20:len(allIP[i])-5], allPort[i][22:len(allPort[i])-5])
			//fmt.Println(str)
			tmpList = append(tmpList, str)
		}
	}
	return tmpList
}
