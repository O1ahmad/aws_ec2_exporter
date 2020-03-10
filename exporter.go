package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	Namespace   = "ec2"
	MetricsPath = "/metrics"
)

type Collector struct {
	instanceTypeVCPUs                                  *prometheus.Desc
}

func NewCollector() *Collector {
	return &Collector{
		instanceTypeVCPUs: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_vcpus"),
			"virtual CPUs offered per instance-type.",
			[]string{"region"},
			nil,
		),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.instanceTypeVCPUs
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	sess := session.Must(session.NewSession())

	svc := ec2.New(sess)
	input := &ec2.DescribeInstanceTypeOfferingsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("state"),
				Values: []*string{
					aws.String("open"),
					aws.String("active"),
					aws.String("closed"),
					aws.String("cancelled"),
					aws.String("failed"),
				},
			},
		},
	}

	result, err := svc.DescribeInstanceTypeOfferings(input)
	if err != nil {
		log.Fatal(err)
	}

	region := os.Getenv("AWS_REGION")

	ch <- prometheus.MustNewConstMetric(
		c.instanceTypeVCPUs,
		prometheus.GaugeValue,
		float64(len(result.instanceTypeVCPUs)),
		region,
	)
}

func init() {
	prometheus.MustRegister(NewCollector())
}

func main() {
	http.Handle(MetricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<html>
			<title>AWS EC2 Exporter</title>
			<body>
			<h1>AWS EC2 Exporter</h1>
			<p><a href='` + MetricsPath + `'>Metrics</a></p>
			</html>`))
		if err != nil {
			log.Printf("Write failed: %v", err)
		}
	})

	log.Print("Starting AWS EC2 Exporter on port 9684...")
	log.Fatal(http.ListenAndServe(":9684", nil))
}
