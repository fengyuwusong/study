package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"
)

func server() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	client := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ResponseHeaderTimeout: time.Second * 60,
			MaxIdleConnsPerHost:   1024,
			MaxIdleConns:          2048,
		},
		Timeout: 30 * time.Second,
	}
	engine.GET("/test", func(context *gin.Context) {
		req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/hello", nil)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("get hello error, err: ", err.Error())
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		//if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
		//	log.Fatal(err)
		//}
		//if err := resp.Body.Close(); err != nil {
		//	log.Fatal("server close body error, err: " + err.Error())
		//}
		context.String(http.StatusOK, fmt.Sprintf("status code: %v", resp.StatusCode))
	})
	pprof.Handler(":8001")
	log.Fatal(engine.Run(":8000"))
}
