package request

type Sample interface {
	GetTimestamp() *Timestamp
}

type QtySample struct {
	Date *Timestamp `json:"date"`
	Qty  float64    `json:"qty"`
}

func (s *QtySample) GetTimestamp() *Timestamp {
	return s.Date
}

type MinMaxAvgSample struct {
	Date *Timestamp `json:"date"`
	Max  float64    `json:"max"`
	Min  float64    `json:"min"`
	Avg  float64    `json:"avg"`
}

func (s *MinMaxAvgSample) GetTimestamp() *Timestamp {
	return s.Date
}

type SleepSample struct {
	Date        *Timestamp `json:"date"`
	Asleep      float64    `json:"asleep"`
	InBed       float64    `json:"inBed"`
	SleepSource string     `json:"sleepSource"`
	InBedSource string     `json:"inBedSource"`
}

func (s *SleepSample) GetTimestamp() *Timestamp {
	return s.Date
}

func newEmptySample(metricType string) Sample {
	switch metricType {
	case MetricTypeQty:
		return &QtySample{}
	case MetricTypeMinMaxAvg:
		return &MinMaxAvgSample{}
	case MetricTypeSleep:
		return &SleepSample{}
	}
	return nil
}
