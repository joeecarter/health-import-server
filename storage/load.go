package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/joeecarter/health-import-server/request"
	"github.com/joeecarter/health-import-server/storage/influxdb"
)

// MetricStore encapsulates a storage backend for the metrics provided by the Auto Export app.
// There is a possibility of the same metrics arriving twice so all MetricStores must handle
// that to avoid storing duplicates.
type MetricStore interface {
	Name() string
	Store(metrics []request.Metric) error
}

type metricStoreLoader func(json.RawMessage) (MetricStore, error)

var metricStoreLoaders = map[string]metricStoreLoader{
	"influxdb": loadInfluxMetricStore,
}

type configType struct {
	Type string `json:"type"`
}

func LoadMetricStores(filename string) ([]MetricStore, error) {
	configs := make([]json.RawMessage, 0)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &configs)
	if err != nil {
		return nil, err
	}

	metricStores := make([]MetricStore, len(configs))
	for i, config := range configs {
		loaderType, err := getConfigType(config)
		if err != nil {
			return nil, err
		}

		loader, ok := metricStoreLoaders[loaderType]
		if !ok {
			logUnknownLoaderType(loaderType, config)
			continue
		}

		metricStore, err := loader(config)
		if err != nil {
			return nil, err
		}

		metricStores[i] = metricStore
	}

	return metricStores, nil
}

func loadInfluxMetricStore(msg json.RawMessage) (MetricStore, error) {
	var config influxdb.InfluxConfig
	if err := json.Unmarshal(msg, &config); err != nil {
		return nil, err
	}
	return influxdb.NewInfluxMetricStore(config), nil
}

func getConfigType(msg json.RawMessage) (string, error) {
	var config configType
	if err := json.Unmarshal(msg, &config); err != nil {
		return "", err
	}
	return config.Type, nil
}

func logUnknownLoaderType(loaderType string, config json.RawMessage) {
	if strings.TrimSpace(loaderType) == "" {
		log.Printf("Encountered an empty loader type. This config will be skipped: %s\n", minifyJson(config))
	} else {
		log.Printf("Encountered an unknown loader type \"%s\". This config will be skipped: %s\n", loaderType, minifyJson(config))
	}
}

// attempts to minfiy the input json swallowing the error if there is one
func minifyJson(b []byte) []byte {
	obj := make(map[string]interface{})
	if err := json.Unmarshal(b, &obj); err != nil {
		return b
	}

	if minified, err := json.Marshal(&obj); err == nil {
		return minified
	}
	return b
}
