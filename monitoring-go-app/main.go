package main

import (
	"monitoring-go-app/logger"
	"monitoring-go-app/metrics"
	"net/http"
	"time"
)

func healthApi(rw http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("Hitting API /health")
	metrics.HealthMetrices.Inc()
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("OK"))
}

func main() {
	logPath := "/var/log/app.log"
	logData := logger.LogData{
		AppName: "golang-app",
		File:    &logPath,
	}
	logData.InitLogger()
	logger.Logger.Info("Starting application ....")

	http.HandleFunc("/health", healthApi)

	go func() {
		for {
			logger.Logger.Info("Testing application ....")
			time.Sleep(time.Minute * 2)
		}
	}()

	// Record metrices
	logger.Logger.Info("Start metrices ...")
	metrics.RegisterMetrices()
	metrics.RecordMetrics()

	logger.Logger.Info("starting server...")
	http.ListenAndServe(":7000", nil) // handlers
	logger.Logger.Info("server is closed")
}
