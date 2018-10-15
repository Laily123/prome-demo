package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

/*
	提供一个 html 页面给prometheus抓取数据
*/

var (
	counter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "counter_demo",
		Help: "just a demo",
	}, []string{"type"})
)

func start() {
	for {
		value := rand.Intn(100)
		counter.With(prometheus.Labels{"type": "demo"}).Add(float64(value))
		time.Sleep(1 * time.Second)
	}
}

func initProme() {
	prometheus.MustRegister(counter)
	go start()
}
