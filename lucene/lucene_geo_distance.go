package lucene

import "strconv"

// GeoDistanceQuery struct
type GeoDistanceQuery struct {
	Type        string  `json:"type"`
	Field       string  `json:"field"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MaxDistance string  `json:"max_distance,omitempty"`
	MinDistance string  `json:"min_distance,omitempty"`
}

// Lucene to match Rule interface
func (v *GeoDistanceQuery) Lucene() bool {
	return true
}

// GeoDistance filter
func GeoDistance(ref string, lat, lon float64, maxDistance interface{}, minDistance ...interface{}) Rule {
	var max, min string
	switch o := maxDistance.(type) {
	case string:
		max = o
	case int8:
		max = strconv.Itoa(int(o)) + "km"
	case int16:
		max = strconv.Itoa(int(o)) + "km"
	case int32:
		max = strconv.Itoa(int(o)) + "km"
	case int64:
		max = strconv.Itoa(int(o)) + "km"
	case uint8:
		max = strconv.Itoa(int(o)) + "km"
	case uint16:
		max = strconv.Itoa(int(o)) + "km"
	case uint32:
		max = strconv.Itoa(int(o)) + "km"
	case uint64:
		max = strconv.Itoa(int(o)) + "km"
	case float32:
		max = strconv.FormatFloat(float64(o), 'f', 6, 64) + "km"
	case float64:
		max = strconv.FormatFloat(o, 'f', 6, 64) + "km"
	default:
		max = "1km"
	}
	if len(minDistance) > 0 {
		switch o := minDistance[0].(type) {
		case string:
			min = o
		case int8:
			min = strconv.Itoa(int(o)) + "km"
		case int16:
			min = strconv.Itoa(int(o)) + "km"
		case int32:
			min = strconv.Itoa(int(o)) + "km"
		case int64:
			min = strconv.Itoa(int(o)) + "km"
		case uint8:
			min = strconv.Itoa(int(o)) + "km"
		case uint16:
			min = strconv.Itoa(int(o)) + "km"
		case uint32:
			min = strconv.Itoa(int(o)) + "km"
		case uint64:
			min = strconv.Itoa(int(o)) + "km"
		case float32:
			min = strconv.FormatFloat(float64(o), 'f', 6, 64) + "km"
		case float64:
			min = strconv.FormatFloat(o, 'f', 6, 64) + "km"
		}
	}
	return &GeoDistanceQuery{
		Type:        "geo_distance",
		Field:       ref,
		Latitude:    lat,
		Longitude:   lon,
		MaxDistance: max,
		MinDistance: min,
	}
}
