package main

import "github.com/prometheus/client_golang/prometheus"

// AddMetrics creates gauge and counter metric vectors based on resource properties
func AddMetrics() (map[string]*prometheus.GaugeVec, map[string]*prometheus.CounterVec) {

	gaugeVecs := make(map[string]*prometheus.GaugeVec)
	counterVecs := make(map[string]*prometheus.CounterVec)

	// instance metrics
	gaugeVecs["totalvCPUs"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "total_vcpus",
			Help:      "Total virtual CPUs capacity provided by an instance-type",
		}, []string{"region", "instance_type", "hypervisor", "bare_metal", "free_tier", "current_gen", "hibernation"})
	gaugeVecs["clockSpeed"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "vcpus_clockspeed",
			Help:      "Clock speed(Ghz) of vCPUs provided by an instance-type",
		}, []string{"region", "instance_type", "hypervisor", "bare_metal", "free_tier", "current_gen", "hibernation"})
	gaugeVecs["totalMem"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "total_mem",
			Help:      "Total memory capacity(GiB) provided by an instance-type",
		}, []string{"region", "instance_type", "hypervisor", "bare_metal", "free_tier", "current_gen", "hibernation"})
	gaugeVecs["totalStorage"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "total_disk",
			Help:      "Total disk storage capacity(GiB) provided by an instance-type",
		}, []string{"region", "instance_type", "hypervisor", "bare_metal", "free_tier", "current_gen", "hibernation"})
	gaugeVecs["ebsOnly"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "ebs_only",
			Help:      "Whether an instance_type *only* supports EBS as its root device volume",
		}, []string{"region", "instance_type", "hypervisor", "bare_metal", "free_tier", "current_gen", "hibernation"})
	gaugeVecs["totalNet"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "total_net",
			Help:      "Total network bandwidth(Gbps) capacity provided by an instance-type",
		}, []string{"region", "instance_type", "hypervisor", "bare_metal", "free_tier", "current_gen", "hibernation"})

	// image metrics
	counterVecs["total_images"] = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "total_images",
			Help:      "Total count of publically available images",
		}, []string{"architecture", "hypervisor", "image_type", "root_device_type", "state", "virtualization_type"})

	// region metrics
	counterVecs["total_regions"] = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "total_regions",
			Help:      "Total count of publically accessible EC2 regions",
		}, []string{"name", "endpoint", "optin_status"})

	// spot instance metrics
	gaugeVecs["spot_price"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "spot_price",
			Help:      "Spot price of EC2 instance types",
		}, []string{"availability_zone", "instance_type", "product_description"})

	return gaugeVecs, counterVecs

}
