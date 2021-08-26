package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mackerelio/go-osstat/memory"
)

type Config struct {
	HTTPAddr     string
	MemThreshold float64
}

//nolint:gochecknoglobals
var defaultConfig Config

func handler(w http.ResponseWriter, r *http.Request) {
	memory, err := memory.Get()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	memPer := float64(memory.Used) / float64(memory.Total) * 100
	// fmt.Println(memPer, memory.Used, memory.Total, memory.Free)
	if memPer > defaultConfig.MemThreshold {
		http.Error(w, "Healthy memory threshold breached", http.StatusServiceUnavailable)
		return
	}

	fmt.Fprintf(w, "Ok!")
}

//nolint:gochecknoinits
func init() {
	flag.StringVar(&defaultConfig.HTTPAddr, "http.addr", ":8080", "HTTP listen address")
	flag.Float64Var(&defaultConfig.MemThreshold, "mem.threshold", 75, "Healthy Memory Threshold")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(defaultConfig.HTTPAddr, nil))
}
