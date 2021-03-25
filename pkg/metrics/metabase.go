package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.etcd.io/bbolt"
)

const metabaseSubsystem = "metabase"

var (
	pageAlloc = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "page_alloc",
		Help:      "pageAlloc",
	})

	nodeDeref = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "node_deref",
		Help:      "nodeDeref",
	})
	times = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "time",
		Help:      "spillTime",
	}, []string{"time"})

	freePageN = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "free_page_n",
		Help:      "freePageN",
	})
	pendingPageN = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "pending_page_n",
		Help:      "pendingPageN",
	})
	freeAlloc = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "free_alloc",
		Help:      "freeAlloc",
	})
	freelistInuse = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "freelist_inuse",
		Help:      "freelistInuse",
	})
)

func init() {
	prometheus.MustRegister(
		pageAlloc,
		nodeDeref,
		times,
		freePageN,
		pendingPageN,
		freeAlloc,
		freelistInuse,
	)
}

func UpdateMetabaseDBStats(s bbolt.Stats) {
	pageAlloc.Set(float64(s.TxStats.PageAlloc))
	nodeDeref.Set(float64(s.TxStats.NodeDeref))
	freePageN.Set(float64(s.FreePageN))
	pendingPageN.Set(float64(s.PendingPageN))
	freeAlloc.Set(float64(s.FreeAlloc))
	freelistInuse.Set(float64(s.FreelistInuse))
	times.With(prometheus.Labels{
		"time": "rebalance",
	}).Set(float64(s.TxStats.RebalanceTime))
	times.With(prometheus.Labels{
		"time": "spill",
	}).Set(float64(s.TxStats.SpillTime))
	times.With(prometheus.Labels{
		"time": "write",
	}).Set(float64(s.TxStats.WriteTime))
}
