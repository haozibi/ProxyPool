package proxy

import (
	"fmt"
	"net/http"

	gg "github.com/haozibi/gglog"
)

var (
	port = "9090"
)

// 以 web 的形式启动代理筛选
func StartProxyByWeb() {
	go startProxy()

	http.HandleFunc("/", index)

	gg.Infoln("listen...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "404 page not found\n")
		return
	}
	fmt.Fprintf(w, "hello world\n")
	return
}
