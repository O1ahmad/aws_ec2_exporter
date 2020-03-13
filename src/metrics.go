package main

import "github.com/prometheus/client_golang/prometheus"

func AddMetrics() (map[string]*prometheus.GaugeVec, map[string]*prometheus.CounterVec) {

  gaugeVecs := make(map[string]*prometheus.GaugeVec)
  counterVecs := make(map[string]*prometheus.CounterVec)

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
  counterVecs["total_images"] = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "total_images",
		Help:      "Total count of publically available images",
	}, []string{"id", "architecture", "hypervisor", "image_type", "root_device_type", "state", "virtualization_type"})

  // region metrics
  counterVecs["total_regions"] = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "total_regions",
		Help:      "Total count of publically accessible EC2 regions",
	}, []string{"name", "endpoint", "optin_status"})

  return gaugeVecs, counterVecs

}
