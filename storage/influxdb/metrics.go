package influxdb

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/joeecarter/health-import-server/request"
)

// TODO: Does this need struct tags??
type InfluxConfig struct {
	Hostname string `json:"hostname"`
	Token    string `json:"token"`
	Org      string `json:"org"`
	Bucket   string `json:"bucket"`
	//Debug    bool   `json:"debug"`
}

type InfluxMetricStore struct {
	writeAPI api.WriteAPIBlocking
	//debugLogging bool
}

func NewInfluxMetricStore(config InfluxConfig) *InfluxMetricStore {
	client := influxdb2.NewClient(config.Hostname, config.Token)
	writeAPI := client.WriteAPIBlocking(config.Org, config.Bucket)
	return &InfluxMetricStore{writeAPI: writeAPI}
}

func (store *InfluxMetricStore) Name() string {
	return "influxdb"
}

func (store *InfluxMetricStore) Store(metrics []request.Metric) error {
	pts := make([]*write.Point, 0)
	for _, metric := range metrics {
		for _, sample := range metric.Samples {
			pts = append(pts, sampleToPoint(metric, sample))
		}
	}

	// NOTE: Influx db handles duplicate points for us so we don't need to worry about that here.
	return store.writeAPI.WritePoint(context.Background(), pts...)
}

func sampleToPoint(metric request.Metric, sample request.Sample) *write.Point {

	p := influxdb2.NewPointWithMeasurement(metric.Name).
		AddField("unit", metric.Unit).
		SetTime(sample.GetTimestamp().ToTime())

	switch v := sample.(type) {

	case *request.QtySample:
		p.AddField("qty", v.Qty)

	case *request.MinMaxAvgSample:
		p.AddField("max", v.Max).
			AddField("min", v.Min).
			AddField("avg", v.Avg)

	case *request.SleepSample:
		p.AddField("asleep", v.Asleep).
			AddField("inBed", v.InBed).
			AddField("sleepSource", v.SleepSource).
			AddField("inBedSource", v.InBedSource)

	default:
		panic(fmt.Sprintf("unexpected Sample type encountered: %T", sample))

	}

	return p
}

//type FlattenedPoint struct {
//	Measurement string
//	Timestamp   time.Time
//	Fields      map[string]interface{}
//}

//func logPoint(pt *write.Point) {
//	flat := FlattenedPoint{}
//	flat.Measurement = pt.Name()
//	flat.Timestamp = pt.Time()

//	fields := make(map[string]interface{})
//	for _, field := range pt.FieldList() {
//		fields[field.Key] = field.Value
//	}
//	flat.Fields = fields

//	b, _ := json.Marshal(flat)
//	fmt.Println(string(b))
//}
