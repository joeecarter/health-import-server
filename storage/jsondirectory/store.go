package jsondirectory

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/joeecarter/health-import-server/request"
)

type JsonDirConfig struct {
	DirectoryPath string `json:"directoryPath"`
}

// JsonDirMetricStore stores metrics in a directory structure.
// Currently this store don't handle duplicates as a new sub directory is created each time.
type JsonDirMetricStore struct {
	config JsonDirConfig
}

func NewJsonDirMetricStore(config JsonDirConfig) *JsonDirMetricStore {
	return &JsonDirMetricStore{config}
}

func (store *JsonDirMetricStore) Store(metrics []request.Metric) error {
	// TODO: Implement
	fmt.Println("TODO: Implement innit")

	dirName := strings.Replace(time.Now().Format(time.RFC3339), ":", "_", -1)
	fmt.Println("dirName =", dirName)
	return nil

	//lookup, err := readLookupFile(store.lookupFilePath())
	//if err != nil {
	//	return 0, err
	//}

	//fmt.Println("yoyo", lookup)

	//for _, metric := range metrics {
	//	jsonSamples, err := json.MarshalIndent(metric.Samples, "", "\t")
	//	if err != nil {
	//		return 0, err
	//	}

	//	ioutil.WriteFile(store.metricFilePath(metric.Name), jsonSamples, 0600)
	//}

	//if err = lookup.save(); err != nil {
	//	return 0, err
	//}

	//return 0, nil
}

func (store *JsonDirMetricStore) metricFilePath(metricName string) string {
	return path.Join(store.config.DirectoryPath, fmt.Sprintf("%s.json", metricName))
}

func (store *JsonDirMetricStore) lookupFilePath() string {
	return path.Join(store.config.DirectoryPath, lookupFileName)
}

//func (store *JsonDirMetricStore) readIndexFile() (indexFile, error) {

//	b, err := ioutil.ReadFile(store.indexFilePath())
//	if err != nil {
//		fmt.Printf("%T: %s\n", err, err.Error())
//	}

//	index := make(indexFile)
//	if err := json.Unmarshal(b, &index); err != nil {
//		return nil, err
//	}
//	return index, nil
//}
