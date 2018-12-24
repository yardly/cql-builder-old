package lucene

// SortGeoDistance struct
type SortGeoDistance struct {
	Type      string  `json:"type"`
	Field     string  `json:"field"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Reverse   bool    `json:"reverse"`
}

// Lucene to match Rule interface
func (v *SortGeoDistance) Lucene() bool {
	return true
}

// ByGeoDistance query
func ByGeoDistance(ref string, lat, lng float64, reverse ...bool) Rule {
	var o bool
	for _, b := range reverse {
		o = b
		break
	}
	return &SortGeoDistance{
		Type:      "geo_distance",
		Field:     ref,
		Latitude:  lat,
		Longitude: lng,
		Reverse:   o,
	}
}
