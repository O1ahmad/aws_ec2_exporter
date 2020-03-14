package main

import (
	"regexp"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) gatherInstanceMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeInstanceTypesOutput, error) {

	var token *string
	var result *ec2.DescribeInstanceTypesOutput
	var err error
	// Describe instance types while token indicates additional paged records
	for ok := true; ok; ok = (token != nil) {
		params := &ec2.DescribeInstanceTypesInput{
			NextToken: token,
		}
		result, err := ec2svc.DescribeInstanceTypes(params)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Debug("DescribeInstanceTypes <RESULT>:", result)

		for _, x := range result.InstanceTypes {
			log.Debug("Data <instance>:", x)
			// total number of vCPUs

			var hypervisor = "unknown"
			if x.Hypervisor != nil {
				hypervisor = *x.Hypervisor
			}
			e.gaugeVecs["totalvCPUs"].With(prometheus.Labels{
				"region":        region,
				"instance_type": *x.InstanceType,
				"hypervisor":    hypervisor,
				"bare_metal":    strconv.FormatBool(*x.BareMetal),
				"free_tier":     strconv.FormatBool(*x.FreeTierEligible),
				"current_gen":   strconv.FormatBool(*x.CurrentGeneration),
				"hibernation":   strconv.FormatBool(*x.HibernationSupported),
			}).Set(float64(*x.VCpuInfo.DefaultVCpus))

			// vCPU maximum supported clockspeed
			if x.ProcessorInfo.SustainedClockSpeedInGhz != nil {
				e.gaugeVecs["clockSpeed"].With(prometheus.Labels{
					"region":        region,
					"instance_type": *x.InstanceType,
					"hypervisor":    hypervisor,
					"bare_metal":    strconv.FormatBool(*x.BareMetal),
					"free_tier":     strconv.FormatBool(*x.FreeTierEligible),
					"current_gen":   strconv.FormatBool(*x.CurrentGeneration),
					"hibernation":   strconv.FormatBool(*x.HibernationSupported),
				}).Set(*x.ProcessorInfo.SustainedClockSpeedInGhz)
			} else {
				e.gaugeVecs["clockSpeed"].With(prometheus.Labels{
					"region":        region,
					"instance_type": *x.InstanceType,
					"hypervisor":    hypervisor,
					"bare_metal":    strconv.FormatBool(*x.BareMetal),
					"free_tier":     strconv.FormatBool(*x.FreeTierEligible),
					"current_gen":   strconv.FormatBool(*x.CurrentGeneration),
					"hibernation":   strconv.FormatBool(*x.HibernationSupported),
				}).Set(-1)
			}
			// total main memory
			e.gaugeVecs["totalMem"].With(prometheus.Labels{
				"region":        region,
				"instance_type": *x.InstanceType,
				"hypervisor":    hypervisor,
				"bare_metal":    strconv.FormatBool(*x.BareMetal),
				"free_tier":     strconv.FormatBool(*x.FreeTierEligible),
				"current_gen":   strconv.FormatBool(*x.CurrentGeneration),
				"hibernation":   strconv.FormatBool(*x.HibernationSupported),
			}).Set(float64(*x.MemoryInfo.SizeInMiB))

			// total disk storage
			var storageSize = 0
			if *x.InstanceStorageSupported {
				storageSize = int(*x.InstanceStorageInfo.TotalSizeInGB)
			}
			e.gaugeVecs["totalStorage"].With(prometheus.Labels{
				"region":        region,
				"instance_type": *x.InstanceType,
				"hypervisor":    hypervisor,
				"bare_metal":    strconv.FormatBool(*x.BareMetal),
				"free_tier":     strconv.FormatBool(*x.FreeTierEligible),
				"current_gen":   strconv.FormatBool(*x.CurrentGeneration),
				"hibernation":   strconv.FormatBool(*x.HibernationSupported),
			}).Set(float64(storageSize))

			// EBS storage ONLY
			var ebsOnly = 0
			if !(*x.InstanceStorageSupported) {
				ebsOnly = 1
			}
			e.gaugeVecs["ebsOnly"].With(prometheus.Labels{
				"region":        region,
				"instance_type": *x.InstanceType,
				"hypervisor":    hypervisor,
				"bare_metal":    strconv.FormatBool(*x.BareMetal),
				"free_tier":     strconv.FormatBool(*x.FreeTierEligible),
				"current_gen":   strconv.FormatBool(*x.CurrentGeneration),
				"hibernation":   strconv.FormatBool(*x.HibernationSupported),
			}).Set(float64(ebsOnly))

			// network bandwidth
			re := regexp.MustCompile(`\d[\d,]*[\.]?[\d{2}]*`)
			netSpeed, _ := strconv.ParseFloat(re.FindString(*x.NetworkInfo.NetworkPerformance), 4)
			e.gaugeVecs["totalNet"].With(prometheus.Labels{
				"region":        region,
				"instance_type": *x.InstanceType,
				"hypervisor":    hypervisor,
				"bare_metal":    strconv.FormatBool(*x.BareMetal),
				"free_tier":     strconv.FormatBool(*x.FreeTierEligible),
				"current_gen":   strconv.FormatBool(*x.CurrentGeneration),
				"hibernation":   strconv.FormatBool(*x.HibernationSupported),
			}).Set(netSpeed)
		}

		// Assign next token for continued processing
		token = result.NextToken
	}

	return result, err

}

func (e *Exporter) gatherImageMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeImagesOutput, error) {

	params := &ec2.DescribeImagesInput{}
	result, err := ec2svc.DescribeImages(params)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Debug("DescribeImages <RESULT>:", result)

	for _, x := range result.Images {
		log.Debug("Data <image>:", x)
		e.counterVecs["total_images"].With(prometheus.Labels{
			"architecture":        *x.Architecture,
			"hypervisor":          *x.Hypervisor,
			"image_type":          *x.ImageType,
			"root_device_type":    *x.RootDeviceType,
			"state":               *x.State,
			"virtualization_type": *x.VirtualizationType,
		}).Inc()
	}

	return result, err

}

func (e *Exporter) gatherRegionMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeRegionsOutput, error) {

	params := &ec2.DescribeRegionsInput{}
	result, err := ec2svc.DescribeRegions(params)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Debug("DescribeRegions <RESULT>:", result)

	for _, x := range result.Regions {
		log.Debug("Data <region>:", x)
		e.counterVecs["total_regions"].With(prometheus.Labels{
			"name":         *x.RegionName,
			"endpoint":     *x.Endpoint,
			"optin_status": *x.OptInStatus,
		}).Inc()
	}

	return result, err

}

func (e *Exporter) gatherSpotMetrics(ch chan<- prometheus.Metric) (*ec2.DescribeSpotPriceHistoryOutput, error) {

	var token *string
	var result *ec2.DescribeSpotPriceHistoryOutput
	var err error
	// Describe historical spot prices for each instance within the past second
	for ok := true; ok; ok = (len(*token) > 0) {
		start := time.Now().Add(-(time.Second * 1))
		end := time.Now()
		params := &ec2.DescribeSpotPriceHistoryInput{
			NextToken: token,
			StartTime: &start,
			EndTime:   &end,
		}
		result, err := ec2svc.DescribeSpotPriceHistory(params)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Debug("Length:", len(result.SpotPriceHistory))
		log.Debug("DescribeSpotHistory <RESULT>:", result)

		for _, x := range result.SpotPriceHistory {
			log.Debug("SpotPrice <instance>:", x)
			var spotPrice, _ = strconv.ParseFloat(*x.SpotPrice, 4)
			e.gaugeVecs["spot_price"].With(prometheus.Labels{
				"availability_zone":   *x.AvailabilityZone,
				"instance_type":       *x.InstanceType,
				"product_description": *x.ProductDescription,
			}).Set(float64(spotPrice))
		}

		// Assign next token for continued processing
		token = result.NextToken
	}

	return result, err

}
