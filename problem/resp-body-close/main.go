package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// 模拟不调用close body异常问题场景 FIXME: 需后续确认原因
// 1. 不进行body读取 及 close 所有环境端口无法复用 端口均为ESTABLISHED状态 内存使用上涨
// 2. 进行body close 云服务器端口上涨到固定数值后可复用，本地windows及虚拟机无法复用 端口为TIME_WAIT
// 3. 进行body读取 所有环境端口可复用
func main() {
	// 启动server
	go func() {
		server()
	}()
	time.Sleep(time.Second)
	client := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          2048,
			MaxIdleConnsPerHost:   1024,
			IdleConnTimeout:       time.Second * 1,
			ResponseHeaderTimeout: time.Second * 60,
		},
		Timeout: 30 * time.Second,
	}

	sum := 0
	for {
		sum++
		fmt.Printf("sum: %d\n", sum)
		req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/test", nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("http request failed, err: %v\n", err)
			time.Sleep(time.Second)
		} else {
			fmt.Printf("http request success, resp.code: %d\n", resp.StatusCode)
		}
		//if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
		//	log.Fatal(err)
		//}
		//if err := resp.Body.Close(); err != nil {
		//	log.Fatal("server close body error, err: " + err.Error())
		//}
	}
}
