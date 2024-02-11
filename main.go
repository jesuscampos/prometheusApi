package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// collectors.WithGoCollectorMemStatsMetricsDisabled()
	// prometheus.Unregister(collectors.NewGoCollector())
	// prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	recordMetrics()

	defaultRegistry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = defaultRegistry
	prometheus.DefaultGatherer = defaultRegistry
	prometheus.MustRegister(opsProcessed)

	http.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
	http.ListenAndServe(":2112", nil)
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_awasome_metric_total",
		Help: "The total number of siths killed",
	})
)
