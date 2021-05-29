package jsondirectory

import (
	"github.com/joeecarter/health-import-server/request"
)

type JsonDirConfig struct {
	DirectoryPath string `json:"directoryPath"`
}

type JsonDirMetricStore struct {
	config JsonDirConfig
}

func NewJsonDirMetricStore(config JsonDirConfig) *JsonDirMetricStore {
	return &JsonDirMetricStore{config}
}

func (store *JsonDirMetricStore) Store(metrics []request.Metric) (int, error) {
	// TODO: Implement
	return 0, nil
}
