package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 创建一个自定义的注册表
	registry := prometheus.NewRegistry()
	// 可选: 添加 process 和 Go 运行时指标到我们自定义的注册表中
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	registry.MustRegister(prometheus.NewGoCollector())

	// 创建一个简单 gauge 指标。
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})

	// 使用我们自定义的注册表注册 gauge
	registry.MustRegister(gauge)

	// 设置 gauge 的值为 39
	gauge.Set(39)

	// 创建Counters 指标
	totalRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of handled HTTP requests.",
	})
	registry.MustRegister(totalRequests)

	totalRequests.Inc()   // +1：计数器增加1.
	totalRequests.Add(23) // +n：计数器增加23.

	// 最终展示
	//# HELP http_requests_total The total number of handled HTTP requests.
	//# TYPE http_requests_total counter
	//http_requests_total 7734

	// Histograms
	requestDurations := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "A histogram of the HTTP request durations in seconds.",
		// Bucket 配置：第一个 bucket 包括所有在 0.05s 内完成的请求，最后一个包括所有在10s内完成的请求。
		// 自动生成线性或者指数增长的 bucket，比如 prometheus.LinearBuckets() 和 prometheus.ExponentialBuckets() 函数。
		Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})
	registry.MustRegister(requestDurations)

	requestDurations.Observe(0.42)

	//// 启动一个计时器
	//timer := prometheus.NewTimer(requestDurations)
	//
	//// [...在应用中处理请求...]
	//
	//// 停止计时器并观察其持续时间，将其放进 requestDurations 的直方图指标中去
	//timer.ObserveDuration()

	// 最终展示
	//# HELP http_request_duration_seconds A histogram of the HTTP request durations in seconds.
	//# TYPE http_request_duration_seconds histogram
	//http_request_duration_seconds_bucket{le="0.05"} 4599
	//http_request_duration_seconds_bucket{le="0.1"} 24128
	//http_request_duration_seconds_bucket{le="0.25"} 45311
	//http_request_duration_seconds_bucket{le="0.5"} 59983
	//http_request_duration_seconds_bucket{le="1"} 60345
	//http_request_duration_seconds_bucket{le="2.5"} 114003
	//http_request_duration_seconds_bucket{le="5"} 201325
	//http_request_duration_seconds_bucket{le="+Inf"} 227420
	//http_request_duration_seconds_sum 88364.234
	//http_request_duration_seconds_count 227420

	// Summaries
	requestSummaryDurations := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "http_request_duration_seconds_summary",
		Help: "A summary of the HTTP request durations in seconds.",
		Objectives: map[float64]float64{
			0.5:  0.05,  // 第50个百分位数，最大绝对误差为0.05。
			0.9:  0.01,  // 第90个百分位数，最大绝对误差为0.01。
			0.99: 0.001, // 第90个百分位数，最大绝对误差为0.001。
		},
	},
	)
	registry.MustRegister(requestSummaryDurations)

	requestDurations.Observe(0.42)

	// 最终展示
	//# HELP http_request_duration_seconds A summary of the HTTP request durations in seconds.
	//# TYPE http_request_duration_seconds summary
	//http_request_duration_seconds{quantile="0.5"} 0.052
	//http_request_duration_seconds{quantile="0.90"} 0.564
	//http_request_duration_seconds{quantile="0.99"} 2.372
	//http_request_duration_seconds_sum 88364.234
	//http_request_duration_seconds_count 227420

	// 标签 新建方法变为：NewXXXXVec
	gaugeVec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_temperature_celsius_vec",
			Help: "The current temperature in degrees Celsius.",
		},
		// 两个标签名称，通过它们来分割指标。
		[]string{"house", "room"},
	)
	registry.MustRegister(gaugeVec)

	// 为 home=ydzs 和 room=living-room 设置指标值
	gaugeVec.WithLabelValues("ydzs", "living-room").Set(27)
	gaugeVec.With(prometheus.Labels{"house": "ydzs", "room": "living-room"}).Set(66)

	// 暴露自定义指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	http.ListenAndServe(":8080", nil)
}
