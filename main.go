package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	prometheus.MustRegister(counter)
	go start()
}

func main() {
	Init()

	http.Handle("/metrics", promhttp.Handler())
	port := "8080"
	fmt.Println("server start: ", port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
