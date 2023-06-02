package receiver

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var metricPrefix string = "logstronaut"

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metricPrefix,
		Name:      "processed_ops_total",
		Help:      "The total number of processed events",
	})
	bytesSaved = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metricPrefix,
		Name:      "db_byte_saved_total",
		Help:      "The total bytes of message saved",
	})
)
