package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joeecarter/health-import-server/storage/influxdb"
)

const INFLUX_HOSTNAME = "INFLUX_HOSTNAME"
const INFLUX_BUCKET = "INFLUX_BUCKET"
const INFLUX_TOKEN = "INFLUX_TOKEN"
const INFLUX_ORG = "INFLUX_ORG"

type metricStoreLoader func(json.RawMessage) (MetricStore, error)

var metricStoreLoaders = map[string]metricStoreLoader{
	"influxdb": loadInfluxMetricStoreFromConfig,
}

type configType struct {
	Type string `json:"type"`
}

func LoadMetricStores(filename string) ([]MetricStore, error) {
	fromConfig, err := LoadMetricStoresFromConfig(filename)
	if err != nil {
		return nil, err
	}

	fromEnvironment, err := LoadMetricStoresFromEnvironment()
	if err != nil {
		return nil, err
	}

	return append(fromConfig, fromEnvironment...), nil
}

func LoadMetricStoresFromConfig(filename string) ([]MetricStore, error) {
	configs := make([]json.RawMessage, 0)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
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

func LoadMetricStoresFromEnvironment() ([]MetricStore, error) {
	metricStore, err := loadInfluxMetricStoreFromEnvironment()
	if err != nil {
		return nil, err
	}

	var metricStores []MetricStore
	if metricStore != nil {
		metricStores = []MetricStore{metricStore}
	}

	return metricStores, nil
}

func loadInfluxMetricStoreFromConfig(msg json.RawMessage) (MetricStore, error) {
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

func loadInfluxMetricStoreFromEnvironment() (MetricStore, error) {
	hostname, hostnameSet := os.LookupEnv(INFLUX_HOSTNAME)
	bucket, bucketSet := os.LookupEnv(INFLUX_BUCKET)
	token, tokenSet := os.LookupEnv(INFLUX_TOKEN)
	org, orgSet := os.LookupEnv(INFLUX_ORG)

	if !hostnameSet && !tokenSet && !orgSet && !bucketSet {
		return nil, nil
	}

	missingVariables := make([]string, 0)
	if !hostnameSet {
		missingVariables = append(missingVariables, INFLUX_HOSTNAME)
	}
	if !bucketSet {
		missingVariables = append(missingVariables, INFLUX_BUCKET)
	}
	if !tokenSet {
		missingVariables = append(missingVariables, INFLUX_TOKEN)
	}
	if !orgSet {
		missingVariables = append(missingVariables, INFLUX_ORG)
	}

	if len(missingVariables) > 0 {
		return nil, missingEnvironmentError{missingVariables}
	}

	//fmt.Printf("hostname = '%s', bucket = '%s', token = '%s', org = '%s'\n", hostname, bucket, token, org)

	config := influxdb.InfluxConfig{
		Hostname: hostname,
		Token:    token,
		Org:      org,
		Bucket:   bucket,
	}
	return influxdb.NewInfluxMetricStore(config), nil
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

type missingEnvironmentError struct {
	missingVariables []string
}

func (err missingEnvironmentError) Error() string {
	return fmt.Sprintf("Missing the following environment variables: [ %s ]", strings.Join(err.missingVariables, ", "))
}
