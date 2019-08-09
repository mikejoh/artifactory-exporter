package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const (
	namespace = "artifactory"
	subsystem = "storage"
)

type storageInfoMetric struct {
	Type  prometheus.ValueType
	Desc  *prometheus.Desc
	Value func(storageInfo storageInfoResponse) float64
}

type StorageInfo struct {
	credentials		  *BasicCredentials
	client            *http.Client
	baseUrl           *url.URL
	up                prometheus.Gauge
	totalScrapes      prometheus.Counter
	jsonParseFailures prometheus.Counter
	metrics           []*storageInfoMetric
}

func NewStorageInfo(httpClient *http.Client, url *url.URL, creds *BasicCredentials) *StorageInfo {

	return &StorageInfo{
		credentials: creds,
		client:  httpClient,
		baseUrl: url,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "up"),
			Help: "Artifactory API endpoint availability.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "total_scrapes"),
			Help: "Current total Artifactory storage API scrapes.",
		}),
		jsonParseFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "json_parse_failures"),
			Help: "Current number of errors while parsing JSON.",
		}),

		metrics: []*storageInfoMetric{
			{
				Type: prometheus.CounterValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "binaries_total"),
					"How many binaries there is in Artifactory",
					nil, nil,
				),
				Value: func(storageInfo storageInfoResponse) float64 {
					return storageInfo.StorageSummary.BinariesSummary.BinariesCount
				},
			},
			{
				Type: prometheus.CounterValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "artifacts_total"),
					"How many artifacts there is in Artifactory",
					nil, nil,
				),
				Value: func(storageInfo storageInfoResponse) float64 {
					return storageInfo.StorageSummary.BinariesSummary.ArtifactsCount
				},
			},
		},
	}
}

// Describe sets the Prometheus metrics descriptions.
func (c *StorageInfo) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		ch <- metric.Desc
	}
	ch <- c.up.Desc()
	ch <- c.totalScrapes.Desc()
	ch <- c.jsonParseFailures.Desc()
}

// Collect collects storage metrics.
func (c *StorageInfo) Collect(ch chan<- prometheus.Metric) {
	c.totalScrapes.Inc()

	defer func() {
		ch <- c.up
		ch <- c.totalScrapes
		ch <- c.jsonParseFailures
	}()

	storageInfoResponse, err := c.fetchAndDecodeStorageInfo()
	if err != nil {
		c.up.Set(0)
		log.Errorf("failed to fetch Storage API data: %s", err)
		return
	}
	c.up.Set(1)

	for _, metric := range c.metrics {
		ch <- prometheus.MustNewConstMetric(
			metric.Desc,
			metric.Type,
			metric.Value(storageInfoResponse),
		)
	}
}

// Fetches the storage info via the Artifactory API
func (c *StorageInfo) fetchAndDecodeStorageInfo() (storageInfoResponse, error) {
	var si storageInfoResponse
	u := *c.baseUrl

	req, err := http.NewRequest("GET", u.String(), nil)
	req.SetBasicAuth(c.credentials.Username, c.credentials.Password)

	res, err := c.client.Do(req)
	if err != nil {
		return si, fmt.Errorf("failed to get storage info from %s://%s:%s%s: %s",
			u.Scheme, u.Hostname(), u.Port(), u.Path, err)
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Errorf("failed to close http.Client", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return si, fmt.Errorf("HTTP Request failed with code %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&si); err != nil {
		c.jsonParseFailures.Inc()
		return si, err
	}

	return si, nil
}
