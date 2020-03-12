package main

import (
  "sync"

  "github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
  mutex     sync.RWMutex
  gaugeVecs map[string]*prometheus.GaugeVec
}

// NewExporter creates the metrics we wish to monitor
func newExporter() *Exporter {

  gaugeVecs := AddMetrics()

  return &Exporter{
	  gaugeVecs: gaugeVecs,
  }
}
