package request

type Workout struct {
	Name                     string         `json:"name"`
	Start                    *Timestamp     `json:"start"`
	End                      *Timestamp     `json:"end"`
	TotalEnergy              QtyUnit        `json:"totalEnergy"`
	ActiveEnergy             QtyUnit        `json:"activeEnergy"`
	AvgHeartRate             QtyUnit        `json:"avgHeartRate"`
	StepCadence              QtyUnit        `json:"stepCadence"`
	Speed                    QtyUnit        `json:"speed"`
	SwimCadence              QtyUnit        `json:"swimCadence"`
	Intensity                QtyUnit        `json:"intensity"`
	Humidity                 QtyUnit        `json:"humidity"`
	TotalSwimmingStrokeCount QtyUnit        `json:"totalSwimmingStrokeCount"`
	FlightsClimbed           QtyUnit        `json:"flightsClimbed"`
	MaxHeartRate             QtyUnit        `json:"maxHeartRate"`
	Distance                 QtyUnit        `json:"distance"`
	StepCount                QtyUnit        `json:"stepCount"`
	Temperature              QtyUnit        `json:"temperature"`
	Elevation                Elevation      `json:"elevation"`
	Route                    []GPSLog       `json:"route"`
	HeartRateData            []HeartRateLog `json:"heartRateData"`
	HeartRateRecovery        []HeartRateLog `json:"heartRateRecovery"`
}

type QtyUnit struct {
	Units string  `json:"units"`
	Qty   float64 `json:"qty"`
}

type GPSLog struct {
	Lat       float64    `json:"lat"`
	Lon       float64    `json:"lon"`
	Altitude  float64    `json:"altitude"`
	Timestamp *Timestamp `json:"timestamp"`
}

type Elevation struct {
	Ascent  float64 `json:"ascent"`
	Descent float64 `json:"descent"`
	Units   string  `json:"units"`
}

type HeartRateLog struct {
	Units string    `json:"units"`
	Date  Timestamp `json:"date"`
	Qty   float64   `json:"qty"`
}
