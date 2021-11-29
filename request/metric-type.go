package request

const MetricTypeQty = "QTY"
const MetricTypeMinMaxAvg = "MINMAXAVG"
const MetricTypeSleep = "SLEEP"
const MetricTypeUnknown = "UNKNOWN"

var metricTypeLookup = map[string]string{
	"active_energy":                      MetricTypeQty,
	"apple_exercise_time":                MetricTypeQty,
	"apple_stand_hour":                   MetricTypeQty,
	"apple_stand_time":                   MetricTypeQty,
	"basal_body_temperature":             MetricTypeQty,
	"basal_energy_burned":                MetricTypeQty,
	"blood_alcohol_content":              MetricTypeQty,
	"blood_glucose":                      MetricTypeUnknown, // TODO: Implement this...
	"blood_oxygen_saturation":            MetricTypeQty,
	"blood_pressure":                     MetricTypeUnknown, // TODO: Implement this...
	"body_fat_percentage":                MetricTypeQty,
	"body_mass_index":                    MetricTypeQty,
	"body_temperature":                   MetricTypeQty,
	"calcium":                            MetricTypeQty,
	"carbohydrates":                      MetricTypeQty,
	"chloride":                           MetricTypeQty,
	"chromium":                           MetricTypeQty,
	"copper":                             MetricTypeQty,
	"cycling_distance":                   MetricTypeQty,
	"dietary_biotin":                     MetricTypeQty,
	"dietary_caffeine":                   MetricTypeQty,
	"dietary_cholesterol":                MetricTypeQty,
	"dietary_energy":                     MetricTypeQty,
	"dietary_sugar":                      MetricTypeQty,
	"dietary_water":                      MetricTypeQty,
	"distance_downhill_snow_sports":      MetricTypeQty,
	"environmental_audio_exposure":       MetricTypeQty,
	"fiber":                              MetricTypeQty,
	"flights_climbed":                    MetricTypeQty,
	"folate":                             MetricTypeQty,
	"forced_expiratory_volume_1":         MetricTypeQty,
	"forced_vital_capacity":              MetricTypeQty,
	"handwashing":                        MetricTypeUnknown, // TODO: Implement this...
	"headphone_audio_exposure":           MetricTypeQty,
	"heart_rate":                         MetricTypeMinMaxAvg,
	"heart_rate_variability":             MetricTypeQty,
	"height":                             MetricTypeQty,
	"high_heart_rate_notifications":      MetricTypeUnknown, // TODO: Implement this...
	"inhaler_usage":                      MetricTypeQty,
	"insulin_delivery":                   MetricTypeQty,
	"iodine":                             MetricTypeQty,
	"iron":                               MetricTypeQty,
	"irregular_heart_rate_notifications": MetricTypeQty,
	"lean_body_mass":                     MetricTypeQty,
	"low_heart_rate_notifications":       MetricTypeQty,
	"magnesium":                          MetricTypeQty,
	"manganese":                          MetricTypeQty,
	"mindful_minutes":                    MetricTypeQty,
	"molybdenum":                         MetricTypeQty,
	"monounsaturated_fat":                MetricTypeQty,
	"niacin":                             MetricTypeQty,
	"number_of_times_fallen":             MetricTypeQty,
	"pantothenic_acid":                   MetricTypeQty,
	"peripheral_perfusion_index":         MetricTypeQty,
	"polyunsaturated_fat":                MetricTypeQty,
	"potassium":                          MetricTypeQty,
	"protein":                            MetricTypeQty,
	"push_count":                         MetricTypeQty,
	"respiratory_rate":                   MetricTypeQty,
	"resting_heart_rate":                 MetricTypeQty,
	"riboflavin":                         MetricTypeQty,
	"saturated_fat":                      MetricTypeQty,
	"selenium":                           MetricTypeQty,
	"sexual_activity":                    MetricTypeUnknown, // TODO: Implement this...
	"six-minute_walking_test_distance":   MetricTypeQty,
	"sleep_analysis":                     MetricTypeSleep, // TODO: Add missing fields
	"sodium":                             MetricTypeQty,
	"stair_speed:_down":                  MetricTypeQty,
	"stair_speed:_up":                    MetricTypeQty,
	"step_count":                         MetricTypeQty,
	"swimming_distance":                  MetricTypeQty,
	"swimming_stroke_count":              MetricTypeQty,
	"thiamin":                            MetricTypeQty,
	"toothbrushing":                      MetricTypeUnknown, // TODO: Implement this...
	"total_fat":                          MetricTypeQty,
	"vo2_max":                            MetricTypeQty,
	"vitamin_a":                          MetricTypeQty,
	"vitamin_b12":                        MetricTypeQty,
	"vitamin_b6":                         MetricTypeQty,
	"vitamin_c":                          MetricTypeQty,
	"vitamin_d":                          MetricTypeQty,
	"vitamin_e":                          MetricTypeQty,
	"vitamin_k":                          MetricTypeQty,
	"waist_circumference":                MetricTypeQty,
	"walking_running_distance":           MetricTypeQty,
	"walking_asymmetry_percentage":       MetricTypeQty,
	"walking_double_support_percentage":  MetricTypeQty,
	"walking_heart_rate_average":         MetricTypeQty,
	"walking_speed":                      MetricTypeQty,
	"walking_step_length":                MetricTypeQty,
	"weight_body_mass":                   MetricTypeQty,
	"wheelchair_distance":                MetricTypeQty,
	"zinc":                               MetricTypeQty,
}

func LookupMetricType(metricName string) string {
	metricType, ok := metricTypeLookup[metricName]
	if ok {
		return metricType
	}
	return MetricTypeUnknown
}
