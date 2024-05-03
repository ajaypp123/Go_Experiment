package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "golang_app_processed_ops_total",
		Help: "The total number of processed events",
	})
)

var (
	HealthMetrices = promauto.NewCounter(prometheus.CounterOpts{
		Name: "golang_app_hits_health_api",
		Help: "The total number of health api",
	})
)

func RegisterMetrices() {
	http.Handle("/metrics", promhttp.Handler())
}

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
