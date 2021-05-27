package request

import (
	"encoding/json"
	"time"
)

const dateLayout = "2006-01-02 15:04:05 -0700"

type SampleTimestamp struct {
	t time.Time
}

func (st *SampleTimestamp) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}
	st.t = t
	return nil
}

func (st *SampleTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.t.Format(dateLayout))
}

func (st *SampleTimestamp) ToTime() time.Time {
	return st.t
}

func (st *SampleTimestamp) String() string {
	return st.t.Format(dateLayout)
}

type Sample interface {
	GetTimestamp() *SampleTimestamp
}

type QtySample struct {
	Date *SampleTimestamp `json:"date"`
	Qty  float64          `json:"qty"`
}

func (s *QtySample) GetTimestamp() *SampleTimestamp {
	return s.Date
}

type MinMaxAvgSample struct {
	Date *SampleTimestamp `json:"date"`
	Max  float64          `json:"max"`
	Min  float64          `json:"min"`
	Avg  float64          `json:"avg"`
}

func (s *MinMaxAvgSample) GetTimestamp() *SampleTimestamp {
	return s.Date
}

type SleepSample struct {
	Date        *SampleTimestamp `json:"date"`
	Asleep      float64          `json:"asleep"`
	InBed       float64          `json:"inBed"`
	SleepSource string           `json:"sleepSource"`
	InBedSource string           `json:"inBedSource"`
}

func (s *SleepSample) GetTimestamp() *SampleTimestamp {
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
