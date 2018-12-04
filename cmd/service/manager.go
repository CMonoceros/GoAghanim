package main

import (
	"github.com/google/tcpproxy"
)

func main() {
	var p tcpproxy.Proxy
	p.AddRoute(":8090", tcpproxy.To("127.0.0.1:8080"))
	p.Run()
}
