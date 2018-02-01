package proxy

func init() {
	addProxy(&proxy{
		name:         "test",
		isAvailable:  true,
		timeInterval: 2,
		getList:      tt,
	})
}

func tt() ([]string, error) {
	return []string{"127.0.0.1:1080"}, nil
}
