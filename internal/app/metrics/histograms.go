package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(RequestDuration)
}

var (
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "bot_request_duration_seconds",
		Help:    "Histogram of the bot request duration in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"command"})
)
