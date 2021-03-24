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
	rebalanceTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "rebalance_time",
		Help:      "rebalanceTime",
	})
	spillTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "spill_time",
		Help:      "spillTime",
	})
	writeTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: metabaseSubsystem,
		Name:      "write_time",
		Help:      "writeTime",
	})
)

func init() {
	prometheus.MustRegister(
		pageAlloc,
		nodeDeref,
		rebalanceTime,
		spillTime,
		writeTime,
	)
}

func UpdateMetabaseDBStats(s bbolt.Stats) {
	pageAlloc.Set(float64(s.TxStats.PageAlloc))
	nodeDeref.Set(float64(s.TxStats.NodeDeref))
	rebalanceTime.Set(float64(s.TxStats.RebalanceTime))
	spillTime.Set(float64(s.TxStats.SpillTime))
	writeTime.Set(float64(s.TxStats.WriteTime))
}
