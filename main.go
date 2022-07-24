package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "working\n")
	log.Printf("%s %s %s", req.Method, req.RequestURI, RemoteIP(req))

}
func index(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			w.Header().Set(name, h)
		}
	}
	os.Setenv("VERSION", "v0.0.1")
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	fmt.Fprintf(w, "headers\n")
	log.Printf("%s %s %s", req.Method, req.RequestURI, RemoteIP(req))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8090", nil)
}
