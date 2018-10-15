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
	testMetrics bool
	configDir   = ""
	port        string
)

func init() {
	flag.BoolVar(&testMetrics, "test", false, "add test metrics")
	flag.StringVar(&configDir, "config-dir", "", "config directory")
	flag.StringVar(&port, "port", "8080", "server port")
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

	if testMetrics {
		go func() {
			mux := http.NewServeMux()
			mux.Handle("/metrics", promhttp.Handler())
			log.Println(":8081/metrics is start")
			http.ListenAndServe(":8081", mux)
		}()
	}

	http.HandleFunc("/add-target", target.AddTargetHandler)
	log.Println("server start: ", port)
	log.Println(http.ListenAndServe(":"+port, nil))
}
