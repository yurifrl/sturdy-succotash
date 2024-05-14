package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	Errors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "app_rrors",
		Help: "Total number of errors",
	})
	Counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "app_counter",
		Help: "Serve as tracer",
	})
)

func init() {
	prometheus.MustRegister(
		Errors,
		Counter,
	)
}
