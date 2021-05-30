package jsondirectory

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/joeecarter/health-import-server/request"
)

const lookupFileName = "LOOKUP.json"

type lookupFile struct {
	sampleRanges map[string]*sampleRange
	filename     string
}

type sampleRange struct {
	Oldest time.Time `json:"oldest"`
	Newest time.Time `json:"newest"`
}

func (file *lookupFile) get(metricName string) *sampleRange {
	return file.sampleRanges[metricName]
}

func (file *lookupFile) update(metricName string, samples []request.Sample) {
	if len(samples) == 0 {
		return
	}

	oldest := samples[0]
	newest := samples[len(samples)-1]

	sampleRange := &sampleRange{
		Oldest: oldest.GetTimestamp().ToTime(),
		Newest: newest.GetTimestamp().ToTime(),
	}

	file.sampleRanges[metricName] = sampleRange
}

func (file *lookupFile) save() error {
	b, err := json.MarshalIndent(file.sampleRanges, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file.filename, b, 0600)
}

func readLookupFile(filename string) (*lookupFile, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return newLookupFile(filename, nil), nil
		}
		return nil, err
	}

	sampleRanges := make(map[string]*sampleRange)
	if err := json.Unmarshal(b, &sampleRanges); err != nil {
		return nil, err
	}
	return newLookupFile(filename, sampleRanges), nil
}

func newLookupFile(filename string, sampleRanges map[string]*sampleRange) *lookupFile {
	if sampleRanges == nil {
		sampleRanges = make(map[string]*sampleRange)
	}
	return &lookupFile{
		sampleRanges: sampleRanges,
		filename:     filename,
	}
}
