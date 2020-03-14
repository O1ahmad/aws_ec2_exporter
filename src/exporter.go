package main

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// Exporter interface for metric registration, collection and forwarding
type Exporter struct {
	mutex       sync.RWMutex
	gaugeVecs   map[string]*prometheus.GaugeVec
	counterVecs map[string]*prometheus.CounterVec
}

// NewExporter creates the metrics we wish to monitor
func newExporter() *Exporter {

	gaugeVecs, counterVecs := AddMetrics()

	return &Exporter{
		gaugeVecs:   gaugeVecs,
		counterVecs: counterVecs,
	}
}
