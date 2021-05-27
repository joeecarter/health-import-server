package main

import "encoding/json"

// Shamelessly stolen from request/json.go

type jsonRequest struct {
	Data jsonData
}

type jsonData struct {
	Metrics []jsonMetric
}

type jsonMetric struct {
	Name  string
	Units string
	Data  []json.RawMessage
}

type metricExample struct {
	Name         string
	OriginalJson string
	MetricFields map[string]interface{}
}

func parseJSONRequest(b []byte) (*jsonRequest, error) {
	req := &jsonRequest{}
	err := json.Unmarshal(b, req)
	return req, err
}

func parseMetricExamples(b []byte) ([]metricExample, error) {
	req, err := parseJSONRequest(b)
	if err != nil {
		return nil, err
	}

	examples := make([]metricExample, len(req.Data.Metrics))
	for i, metric := range req.Data.Metrics {
		var metricFields map[string]interface{}
		var originalJson = ""
		if len(metric.Data) > 0 {
			metricFields = make(map[string]interface{})
			originalJson = string(metric.Data[0])
			err = json.Unmarshal(metric.Data[0], &metricFields)
			if err != nil {
				return nil, err
			}
		}

		examples[i] = metricExample{
			Name:         metric.Name,
			OriginalJson: originalJson,
			MetricFields: metricFields,
		}
	}

	return examples, nil
}
