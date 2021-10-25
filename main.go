package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mackerelio/go-osstat/memory"
)

type Config struct {
	HTTPAddr     string
	MemThreshold float64
	Cli          bool
	Unhealthy    bool
}

//nolint:gochecknoglobals
var defaultConfig Config

func isHealthy() (bool, error) {
	mem, err := memory.Get()
	if err != nil {
		return true, err
	}

	memPer := float64(mem.Used) / float64(mem.Total) * 100

	log.Println("Mem Per%", memPer, "Mem Used", mem.Used, "Mem Total", mem.Total, "Mem Free", mem.Free)

	if memPer > defaultConfig.MemThreshold {
		return false, nil
	}

	return true, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	h, _ := isHealthy()

	if !h {
		http.Error(w, "Healthy memory threshold breached", http.StatusServiceUnavailable)
	}

	fmt.Fprintf(w, "Ok!")
}

//nolint:gochecknoinits
func init() {
	flag.StringVar(&defaultConfig.HTTPAddr, "http.addr", ":8080", "HTTP listen address")
	flag.Float64Var(&defaultConfig.MemThreshold, "mem.threshold", 75, "Healthy Memory Threshold")
	flag.BoolVar(&defaultConfig.Unhealthy, "unhealthy", false, "Exit with code 0 on unhealthy")
	flag.BoolVar(&defaultConfig.Cli, "cli", false, "Cli Only")
}

func main() {
	flag.Parse()

	exCode := 0

	if defaultConfig.Unhealthy {
		exCode = 1
	}

	if defaultConfig.Cli {
		h, _ := isHealthy()

		if h {
			os.Exit(exCode)
		}

		os.Exit((exCode + 1) % 2)
	}

	log.Println("Starting HTTP server on", defaultConfig.HTTPAddr)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(defaultConfig.HTTPAddr, nil))
}
