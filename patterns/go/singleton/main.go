package main

import "sync"

func Metrics() HandlerFunc {
	metricMap := make(map[string]m3.Metric) // key: handler name
	pathMap := make(map[string][]string)    // key: handler name, value: paths
	once := sync.Once{}                     // protect maps init

	return func(c *Context) {
		// why init in handler chain not earlier: routes haven't been registered into the engine when the middleware func created.
		once.Do(func() {
			for _, r := range c.engine.Routes() {
				metricMap[r.Handler] = cli.NewMetric(r.Handler+".calledby", tagMethod, tagURI, tagErrCode, tagFromCluster, tagToCluster, tagFrom, tagEnv)
				pathMap[r.Handler] = append(pathMap[r.Handler], r.Path)
			}
		})

		c.Next()
	}
}
