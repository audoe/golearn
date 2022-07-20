package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "hello\n")
	Log(req)

}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func Log(r *http.Request) {
	fmt.Println(time.Now().Format(time.RFC1123), "访问地址:", r.URL, RemoteIP(r))
}

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "200\n")
	Log(req)

}
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			w.Header().Set(name, h)
		}
	}

	w.Header().Set("VERSION", os.Getenv("VERSION"))
	fmt.Fprintf(w, "hello\n")
	Log(req)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/healthz", healthz)

	http.ListenAndServe(":8090", nil)
}
