package main

import (
	"net/http"
	"net/url"

	"github.com/mikejoh/artifactory-exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		bind 	= kingpin.Flag("web.listen-address", "Address:Port to listen on for web interface and telemetry.").Default(":9627").String()
		metrics = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
		apiUrl	= kingpin.Flag("web.artifactory-api-url", "Artifactory REST API URL.").Default("http://localhost:8081/artifactory/api/storageinfo").String()
		user	= kingpin.Flag("web.username", "Artifactory API user.").Default("admin").Short('u').String()
		pass	= kingpin.Flag("web.password", "Artifactory API user password.").Default("password").Short('p').String()
	)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	u, err := url.Parse(*apiUrl)
	if err != nil {
		log.Fatalf("failed parsing url", err)
	}

	bc := &collector.BasicCredentials{
		Username: *user,
		Password: *pass,
	}

	httpClient := &http.Client{}

	prometheus.MustRegister(collector.NewStorageInfo(httpClient, u, bc))

	mux := http.DefaultServeMux
	mux.Handle(*metrics, promhttp.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err = w.Write([]byte(`<html>
			<head><title>Artifactory Exporter</title></head>
			<body>
			<h1>Artifactory Exporter</h1>
			<p><a href="` + *metrics + `">Metrics</a></p>
			</body>
			</html>`))
		if err != nil {
			log.Errorf("failed handling writer", err)
		}
	})

	server := &http.Server{
		Addr: *bind,
		Handler: mux,
	}

	log.Infof("starting Artifactory exporter on: %s", *bind)
	log.Fatal(server.ListenAndServe())
}