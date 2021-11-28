package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 实现net/http ListenAndServe Handler 接口
type Engine struct {
}

func (e Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		fmt.Fprintf(writer, "URL.Path: %v", request.URL.Path)
	case "/hello":
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(writer, "404 NOT FOUNE:%s\n", request.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatalln(http.ListenAndServe(":9999", engine))
}
