package proxy

import "regexp"

func init() {
	addProxy(&proxy{
		name:         "66ip.cn",
		isAvailable:  true,
		timeInterval: 60,
		getList:      get66IP,
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
