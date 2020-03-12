package main

import "github.com/prometheus/client_golang/prometheus"

func AddMetrics() map[string]*prometheus.GaugeVec {

  gaugeVecs := make(map[string]*prometheus.GaugeVec)

  // instance metrics
  gaugeVecs["totalvCPUs"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "total_vcpus",
		Help:      "Total virtual CPUs capacity provided by an instance-type",
	}, []string{"region", "instance_type"})
  gaugeVecs["clockSpeed"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "vcpus_clockspeed",
		Help:      "Clock speed(Ghz) of vCPUs provided by an instance-type",
	}, []string{"region", "instance_type"})
  gaugeVecs["totalMem"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "total_mem",
		Help:      "Total memory capacity(GiB) provided by an instance-type",
	}, []string{"region", "instance_type"})
  gaugeVecs["totalStorage"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "total_disk",
		Help:      "Total disk storage capacity(GiB) provided by an instance-type",
	}, []string{"region", "instance_type"})
  gaugeVecs["ebsOnly"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "ebs_only",
		Help:      "Whether an instance_type *only* supports EBS as its root device volume",
	}, []string{"region", "instance_type"})
  gaugeVecs["totalNet"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "total_net",
		Help:      "Total network bandwidth(Gbps) capacity provided by an instance-type",
	}, []string{"region", "instance_type"})

  // image metrics
  gaugeVecs["imageState"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "image_state",
		Help:      "The current state of the AMI (available | pending | failed)",
	}, []string{"name", "id", "hypervisor"})
  gaugeVecs["imageSize"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "image_size",
		Help:      "The total size of the image(MiB)",
	}, []string{"name", "id", "hypervisor"})

  // region metrics
  gaugeVecs["regionStatus"] = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "region_status",
		Help:      "The region opt-in status (opt-in-not-required | opted-in | not-opted-in)",
	}, []string{"name", "endpoint"})

  return gaugeVecs

}
