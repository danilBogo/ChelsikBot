package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(TotalRequestsCounter)
	prometheus.MustRegister(TotalCommandCounter)
	prometheus.MustRegister(TotalUserCommandCounter)
	prometheus.MustRegister(SuccessCommandCounter)
	prometheus.MustRegister(SuccessUserCommandCounter)
}

var (
	TotalRequestsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "bot_requests_total",
		Help: "Total requests number",
	})

	TotalCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_total_command",
		Help: "Number of total command requests",
	}, []string{"command"})

	TotalUserCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_total_user_command",
		Help: "Number of total user command requests",
	}, []string{"username", "command"})

	SuccessCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_success_command",
		Help: "Number of success command requests",
	}, []string{"command"})

	SuccessUserCommandCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_requests_success_user_command",
		Help: "Number of success user command requests",
	}, []string{"username", "command"})
)
