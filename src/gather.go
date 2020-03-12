package main

import (
    "strings"
    "strconv"
	log "github.com/Sirupsen/logrus"

    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) gatherInstanceMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeInstanceTypesOutput, error) {

    ec2svc := ec2.New(session.New(&aws.Config{Region: aws.String(region),}))
	params := &ec2.DescribeInstanceTypesInput{
	}
	result, err := ec2svc.DescribeInstanceTypes(params)
	if err != nil {
	  log.Fatal(err.Error())
	}

	log.Debugf("Data Captured", result)

	for _, x := range result.InstanceTypes {
      log.Debugf("Data Captured", x)
        e.gaugeVecs["totalvCPUs"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(*x.VCpuInfo.DefaultVCpus))
        e.gaugeVecs["clockSpeed"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(*x.ProcessorInfo.SustainedClockSpeedInGhz)
		e.gaugeVecs["totalMem"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(*x.MemoryInfo.SizeInMiB))
        var storage_size = 0
        if *x.InstanceStorageSupported {
          storage_size = int(*x.InstanceStorageInfo.TotalSizeInGB)
        } else {
          storage_size = 0
        }
		e.gaugeVecs["totalStorage"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(storage_size))
        var ebs_only = 0
        if !(*x.InstanceStorageSupported) {
          ebs_only = 1
        }
		e.gaugeVecs["ebsOnly"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(ebs_only))
        if v, err := strconv.ParseFloat(strings.SplitN(*x.NetworkInfo.NetworkPerformance, " ", 2)[0], 4); err == nil {
		  e.gaugeVecs["totalNet"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(v)
        } else {
		  e.gaugeVecs["totalNet"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(-1)
        }
    }

	return result, err

}

func (e *Exporter) gatherImageMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeImagesOutput, error) {

    ec2svc := ec2.New(session.New(&aws.Config{Region: aws.String(region),}))
	params := &ec2.DescribeImagesInput{
	}
	result, err := ec2svc.DescribeImages(params)
	if err != nil {
	  log.Fatal(err.Error())
	}

	log.Debugf("Data Captured", result)

	for _, x := range result.Images {
      log.Debugf("Data Captured", x)
		//e.gaugeVecs["imageState"].With(prometheus.Labels{"name": x.Name, "id": x.Id, "hypervisor": x.Hypervisor}).Set(x.State)
		//e.gaugeVecs["imageSize"].With(prometheus.Labels{"name": x.Name, "id": x.Id, "hypervisor": x.hypervisor}).Set(x.Size)
	}

	return result, err

}

func (e *Exporter) gatherRegionMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeRegionsOutput, error) {

    ec2svc := ec2.New(session.New(&aws.Config{Region: aws.String(region),}))
	params := &ec2.DescribeRegionsInput{
	}
	result, err := ec2svc.DescribeRegions(params)
	if err != nil {
	  log.Fatal(err.Error())
	}

	log.Debugf("Data Captured", result)

	for _, x := range result.Regions {
      log.Debugf("Data Captured", x)
      //e.gaugeVecs["regionStatus"].With(prometheus.Labels{"name": x.Name, "endpoint": x.endpoint}).Set(x.Status)
	}

	return result, err

}
