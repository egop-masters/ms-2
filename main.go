package main

import (
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

// Define a counter metric
var (
    requestCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Number of HTTP requests received.",
        },
        []string{"path"},
    )
)

func init() {
    // Register the counter with Prometheus's default registry
    prometheus.MustRegister(requestCount)
}

func main() {
    r := gin.Default()

    // Define the GET endpoint
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))
    r.GET("/test-endpoint", func(c *gin.Context) {
        path := c.FullPath()
        requestCount.WithLabelValues(path).Inc()
        c.String(http.StatusOK, "Hello from ms-2!")
    })

    // Run the server
    r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}

