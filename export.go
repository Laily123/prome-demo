package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

/*
	提供一个 html 页面给prometheus抓取数据
*/

var (
	counter = prometheus.NewCounterVec("")
)
