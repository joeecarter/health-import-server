package request

import (
	"encoding/json"
	"log"
)

var LogUnknownMetrics = false

type APIExportRequest struct {
	Metrics []Metric `json:"name"`
}

func (req *APIExportRequest) TotalSamples() int {
	total := 0
	for _, metric := range req.Metrics {
		total += len(metric.Samples)
	}
	return total
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
		metric, err := p.ParseMetric(jsonMetric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return &APIExportRequest{
		Metrics: metrics,
	}, nil
}

func (p *Parser) ParseMetric(jm jsonMetric) (Metric, error) {
	samples, err := p.ParseSamples(jm.Name, jm.Data)
	if err != nil {
		return Metric{}, err
	}

	return Metric{
		Name:    jm.Name,
		Unit:    jm.Units,
		Samples: samples,
	}, nil
}

func (p *Parser) ParseSamples(metricName string, rawSamples []json.RawMessage) ([]Sample, error) {
	metricType := LookupMetricType(metricName)
	if metricType == MetricTypeUnknown {
		p.logUnknownMetric(metricName, rawSamples)
		return nil, nil
	}
	return p.unmarshalSamples(metricType, rawSamples)
}

func (p *Parser) unmarshalSamples(metricType string, rawSamples []json.RawMessage) ([]Sample, error) {
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

func (p *Parser) logUnknownMetric(metricName string, rawSamples []json.RawMessage) {
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
