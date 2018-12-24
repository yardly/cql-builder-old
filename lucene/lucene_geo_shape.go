package lucene

import "fmt"

// GeoShapeQuery struct
type GeoShapeQuery struct {
	Type      string `json:"type"`
	Field     string `json:"field"`
	Shape     string `json:"shape"`
	Operation string `json:"operation"`
}

// Lucene to match Rule interface
func (v *GeoShapeQuery) Lucene() bool {
	return true
}

// GeoShape filter
func GeoShape(ref, shape string, operations ...string) Rule {
	v := "contains"
	for _, operation := range operations {
		v = operation
	}
	return &GeoShapeQuery{
		Type:      "geo_shape",
		Field:     ref,
		Shape:     shape,
		Operation: v,
	}
}

func GeoShapePoint(ref string, lat, lng float64, operations ...string) Rule {
	return GeoShape(ref, fmt.Sprintf("POINT(%v %v)", lng, lat), operations...)
}
