package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"prome-demo/target"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	test      bool
	configDir = ""
)

func init() {
	flag.BoolVar(&test, "test", false, "add test metrics")
	flag.StringVar(&configDir, "config-dir", "", "config directory")
	flag.Parse()
	if configDir == "" {
		log.Println("please add config dir")
		os.Exit(0)
	}
	target.ConfigDir = configDir
}

func initProme() {
	prometheus.MustRegister(counter)
	go start()
}

func main() {
	initProme()

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", mux)
	}()

	http.HandleFunc("/add-target", target.AddTargetHandler)
	port := "8080"
	log.Println("server start: ", port)
	log.Println(http.ListenAndServe(":"+port, nil))
}
