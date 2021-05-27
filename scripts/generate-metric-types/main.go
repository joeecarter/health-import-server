package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("request.json")
	if err != nil {
		panic("Failed to read file request.json err = " + err.Error())
	}

	examples, err := parseMetricExamples(b)
	if err != nil {
		panic("Failed parse metric examples err = " + err.Error())
	}

	metricNames := make([]string, len(examples))
	for i, example := range examples {
		metricNames[i] = example.Name
	}

	metricTypes := inferMetricTypes(examples)
	//printJson(metricTypes)

	code := generateCode(metricNames, metricTypes)
	fmt.Println(code)
}

func generateCode(metricNames []string, metricTypes map[string]string) string {
	b := strings.Builder{}
	longest := longestString(metricNames)

	b.WriteString("var metricTypeLookup = map[string]string{\n")

	for _, metricName := range metricNames {
		metricType := metricTypes[metricName]
		padding := strings.Repeat(" ", 1+longest-len(metricName))
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", metricName, padding, metricType))
	}

	b.WriteString("}")

	return b.String()
}

// returns map[MetricName]MatricType
func inferMetricTypes(examples []metricExample) map[string]string {
	var results = make(map[string]string)
	for _, example := range examples {
		results[example.Name] = interMetricType(example)
	}
	return results
}

func interMetricType(example metricExample) string {
	// Nothing to work with - cannot infer
	if example.MetricFields == nil {
		return "MetricTypeUnknown"
	}

	if isQtyMetricType(example.MetricFields) {
		return "MetricTypeQty"
	}

	if isMinMaxAvgMetricType(example.MetricFields) {
		return "MetricTypeMinMaxAvg"
	}

	if isSleepMetricType(example.MetricFields) {
		return "MetricTypeSleep"
	}

	fmt.Printf("I'd recommend you investigate metric %s - %s\n", example.Name, example.OriginalJson)

	return "MetricTypeUnknown"
}

func isQtyMetricType(metricFields map[string]interface{}) bool {
	if len(metricFields) != 2 {
		return false
	}

	if !fieldExists(metricFields, "date") {
		return false
	}
	if !fieldExists(metricFields, "qty") {
		return false
	}

	if !isString(metricFields["date"]) {
		return false
	}
	if !isFloat(metricFields["qty"]) {
		return false
	}

	return true
}

func isMinMaxAvgMetricType(metricFields map[string]interface{}) bool {
	if len(metricFields) != 4 {
		return false
	}

	if !fieldExists(metricFields, "date") {
		return false
	}
	if !fieldExists(metricFields, "Max") {
		return false
	}
	if !fieldExists(metricFields, "Min") {
		return false
	}
	if !fieldExists(metricFields, "Avg") {
		return false
	}

	if !isString(metricFields["date"]) {
		return false
	}
	if !isFloat(metricFields["Max"]) {
		return false
	}
	if !isFloat(metricFields["Min"]) {
		return false
	}
	if !isFloat(metricFields["Avg"]) {
		return false
	}

	return true
}

func isSleepMetricType(metricFields map[string]interface{}) bool {
	if len(metricFields) != 5 {
		return false
	}

	if !fieldExists(metricFields, "date") {
		return false
	}
	if !fieldExists(metricFields, "asleep") {
		return false
	}
	if !fieldExists(metricFields, "inBed") {
		return false
	}
	if !fieldExists(metricFields, "sleepSource") {
		return false
	}
	if !fieldExists(metricFields, "inBedSource") {
		return false
	}

	if !isString(metricFields["date"]) {
		return false
	}
	if !isFloat(metricFields["asleep"]) {
		return false
	}
	if !isFloat(metricFields["inBed"]) {
		return false
	}
	if !isString(metricFields["sleepSource"]) {
		return false
	}
	if !isString(metricFields["inBedSource"]) {
		return false
	}

	return true
}

func fieldExists(metricFields map[string]interface{}, fieldName string) bool {
	_, ok := metricFields[fieldName]
	return ok
}

func isString(v interface{}) bool {
	_, ok := v.(string)
	return ok
}

func isFloat(v interface{}) bool {
	_, ok := v.(float64)
	return ok
}

func longestString(ss []string) int {
	longest := 0
	for _, s := range ss {
		len := len(s)
		if len > longest {
			longest = len
		}
	}
	return longest
}

func printJson(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println(string(b))
}
