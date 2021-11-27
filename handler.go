package server

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joeecarter/health-import-server/request"
)

type ImportHandler struct {
	MetricStores []MetricStore
}

func NewImportHandler(metricStores []MetricStore) *ImportHandler {
	return &ImportHandler{metricStores}
}

func (handler *ImportHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	msg, err := handler.handle(req)
	if err == nil {
		wr.WriteHeader(200)
		wr.Write([]byte(msg + "\n"))
	} else {
		wr.WriteHeader(500)
		wr.Write([]byte("ERROR: " + err.Error() + "\n"))
	}
}

func (handler *ImportHandler) handle(req *http.Request) (string, error) {
	log.Printf("Received request with User-Agent: '%s'\n", req.Header.Get("User-Agent"))

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic("Failed to get body, err =" + err.Error())
	}

	export, err := request.Parse(b)
	if err != nil {
		return "", err
	}

	populatedMetrics := export.PopulatedMetrics()
	log.Printf("Total metrics: %d (%d populated) Total samples %d\n", len(export.Metrics), len(populatedMetrics), export.TotalSamples())

	// TODO: May want to not fail fast here? run all metric stores before erroring to avoid data loss?
	for _, metricStore := range handler.MetricStores {
		log.Printf("Starting upload to metric store \"%s\".", metricStore.Name())
		if err = metricStore.Store(populatedMetrics); err != nil {
			log.Printf("Failed upload to metric store \"%s\" with error: %s.", metricStore.Name(), err.Error())
			return "", err
		}
		log.Printf("Finished upload to metric store \"%s\".", metricStore.Name())
	}

	return "Processed request.", nil
}
