package proxy

import "regexp"

func init() {
	addProxy(&proxy{
		name:         "66ip.cn", // 名称
		isAvailable:  true,      // 是否启用
		timeInterval: 60,        // 每次启动间隔，以秒为单位
		getList:      get66IP,   // 具体方法
	})
}

var (
	tmpIpUri = "http://m.66ip.cn/mo.php?sxb=&tqsl=100&port=&export=&ktip=&sxa=&submit=%CC%E1++%C8%A1&textare"
)

func get66IP() ([]string, error) {
	resp, _, err := HttpFunc(tmpIpUri, "", "GET")
	if err != nil {
		return []string{}, err
	}
	reg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}`)
	return reg.FindAllString(string(resp), -1), nil
}
