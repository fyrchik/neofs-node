package meta

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister()
}
