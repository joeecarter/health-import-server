package request

import (
	"encoding/json"
	"log"
)

var LogUnknownMetrics = false

type APIExportRequest struct {
	Metrics          []Metric `json:"metrics"`
	populatedMetrics []Metric
}

// TotalSamples returns the total number of samples across all metrics
func (req *APIExportRequest) TotalSamples() int {
	total := 0
	for _, metric := range req.PopulatedMetrics() {
		total += len(metric.Samples)
	}
	return total
}

// PopulatedMetrics returns only the metrics which have > 0 samples.
func (req *APIExportRequest) PopulatedMetrics() []Metric {
	if req.populatedMetrics != nil {
		return req.populatedMetrics
	}

	req.populatedMetrics = make([]Metric, 0)
	for _, metric := range req.Metrics {
		if len(metric.Samples) == 0 {
			continue
		}
		req.populatedMetrics = append(req.populatedMetrics, metric)
	}
	return req.populatedMetrics
}

type Metric struct {
	Name    string   `json:"name"`
	Unit    string   `json:"unit"`
	Samples []Sample `json:"samples"`
}

func Parse(b []byte) (*APIExportRequest, error) {
	p := NewParser(b)
	return p.Parse()
}

type Parser struct {
	b []byte
}

func NewParser(b []byte) *Parser {
	return &Parser{b}
}

func (p *Parser) Parse() (*APIExportRequest, error) {
	j, err := parseJSONRequest(p.b)
	if err != nil {
		return nil, err
	}

	metrics := make([]Metric, 0)
	for _, jsonMetric := range j.Data.Metrics {
		metric, err := p.parseMetric(jsonMetric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return &APIExportRequest{
		Metrics: metrics,
	}, nil
}

func (p *Parser) parseMetric(jm jsonMetric) (Metric, error) {
	samples, err := ParseSamples(jm.Name, jm.Data)
	if err != nil {
		return Metric{}, err
	}

	return Metric{
		Name:    jm.Name,
		Unit:    jm.Units,
		Samples: samples,
	}, nil
}

func ParseSamples(metricName string, rawSamples []json.RawMessage) ([]Sample, error) {
	metricType := LookupMetricType(metricName)
	if metricType == MetricTypeUnknown {
		logUnknownMetric(metricName, rawSamples)
		return nil, nil
	}
	return unmarshalSamples(metricType, rawSamples)
}

func unmarshalSamples(metricType string, rawSamples []json.RawMessage) ([]Sample, error) {
	samples := make([]Sample, len(rawSamples))
	for i, rawSample := range rawSamples {
		sample := newEmptySample(metricType)
		err := json.Unmarshal(rawSample, sample)
		if err != nil {
			return nil, err
		}

		samples[i] = sample
	}

	return samples, nil
}

func logUnknownMetric(metricName string, rawSamples []json.RawMessage) {
	if !LogUnknownMetrics {
		return
	}

	var example json.RawMessage = nil
	if len(rawSamples) > 0 {
		example = rawSamples[0]
	}

	exampleString := "null"
	if example != nil {
		exampleString = string(example)
	}

	log.Printf("Encountered unknown metric '%s' example = %s\n\n", metricName, exampleString)
}
