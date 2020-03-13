package main

import (
    "strconv"
    "regexp"
	log "github.com/Sirupsen/logrus"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) gatherInstanceMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeInstanceTypesOutput, error) {

	params := &ec2.DescribeInstanceTypesInput{
	}
	result, err := ec2svc.DescribeInstanceTypes(params)
	if err != nil {
	  log.Fatal(err.Error())
	}

    log.Debug("Data Captured:", result)

	for _, x := range result.InstanceTypes {
      log.Debug("Data Captured", x)
      // total number of vCPUs
      e.gaugeVecs["totalvCPUs"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(*x.VCpuInfo.DefaultVCpus))

      // vCPU maximum supported clockspeed
      if x.ProcessorInfo.SustainedClockSpeedInGhz != nil {
           e.gaugeVecs["clockSpeed"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(*x.ProcessorInfo.SustainedClockSpeedInGhz) } else {
           e.gaugeVecs["clockSpeed"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(-1)
      }
      // total main memory
      e.gaugeVecs["totalMem"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(*x.MemoryInfo.SizeInMiB))

      // total disk storage
      var storage_size = 0
      if *x.InstanceStorageSupported {
        storage_size = int(*x.InstanceStorageInfo.TotalSizeInGB)
      }
      e.gaugeVecs["totalStorage"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(storage_size))

      // EBS storage ONLY
      var ebs_only = 0
      if !(*x.InstanceStorageSupported) {
        ebs_only = 1
      }
      e.gaugeVecs["ebsOnly"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(float64(ebs_only))

      // network bandwith
      re := regexp.MustCompile(`\d[\d,]*[\.]?[\d{2}]*`)
      net_speed, _ := strconv.ParseFloat(re.FindString(*x.NetworkInfo.NetworkPerformance), 4)
      e.gaugeVecs["totalNet"].With(prometheus.Labels{"region": region, "instance_type": *x.InstanceType}).Set(net_speed)
    }

	return result, err

}

func (e *Exporter) gatherImageMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeImagesOutput, error) {

	params := &ec2.DescribeImagesInput{
	}
	result, err := ec2svc.DescribeImages(params)
	if err != nil {
	  log.Fatal(err.Error())
	}

	for _, x := range result.Images {
      log.Debug("Data Captured:", x)
		//e.gaugeVecs["imageState"].With(prometheus.Labels{"name": x.Name, "id": x.Id, "hypervisor": x.Hypervisor}).Set(x.State)
		//e.gaugeVecs["imageSize"].With(prometheus.Labels{"name": x.Name, "id": x.Id, "hypervisor": x.hypervisor}).Set(x.Size)
	}

	return result, err

}

func (e *Exporter) gatherRegionMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeRegionsOutput, error) {

	params := &ec2.DescribeRegionsInput{
	}
	result, err := ec2svc.DescribeRegions(params)
	if err != nil {
	  log.Fatal(err.Error())
	}

    log.Debug("Data Captured:", result)

	for _, x := range result.Regions {
      log.Debug("Data Captured", x)
      //e.gaugeVecs["regionStatus"].With(prometheus.Labels{"name": x.Name, "endpoint": x.endpoint}).Set(x.Status)
	}

	return result, err

}
