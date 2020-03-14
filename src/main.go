package main

import (
	"net/http"

	c "github.com/0x0I/aws_ec2_exporter/src/config"
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "aws_ec2" // Used to prepand Prometheus metrics
)

// Runtime variables
var (
	metricsPath = c.GetEnv("METRICS_PATH", "/metrics") // Path under which to expose metrics
	listenPort  = c.GetEnv("LISTEN_PORT", ":9686")     // Port on which to expose metrics
	logLevel    = c.GetEnv("LOG_LEVEL", "info")
	region      = c.GetEnv("REGION", "us-east-1")

	ec2svc = ec2.New(session.New(&aws.Config{Region: aws.String(region)})) // EC2 API session/service client
)

func main() {
	c.CheckConfig()

	setLogLevel(logLevel)
	log.Info("Starting Prometheus AWS EC2 Exporter")

	exporter := newExporter()
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
      <head><title>AWS EC2 Exporter</title></head>
      <body>
      <h1>AWS EC2 Exporter</h1>
      <p><a href=` + metricsPath + `>Metrics</a></p>
      </body>
      </html>`))
	})

	log.Printf("Starting Server on port %s and path %s", listenPort, metricsPath)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
