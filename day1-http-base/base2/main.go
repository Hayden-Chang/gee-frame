package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engin结构体
type Engine struct{}

// 实现ServeHTTP方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
// main方法启动
func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8888", engine))
}