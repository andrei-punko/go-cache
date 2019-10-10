package datatype

import "time"

// DataType represents cache item with Value stored inside, Ttl and DeathTime fields.
type DataType struct {
	Value     interface{}   `json:"value"`
	Ttl       time.Duration `json:"ttl"`
	DeathTime time.Time     `json:"deathTime"`
}

// NewString creates DataType item with string value inside.
// Its DeathTime = (current time) + (provided TTL).
func NewString(value string, ttl time.Duration) DataType {
	return DataType{value, ttl, time.Now().Add(ttl)}
}

// NewList creates DataType item with list value.
// Its DeathTime = (current time) + (provided TTL).
func NewList(value []interface{}, ttl time.Duration) DataType {
	return DataType{value, ttl, time.Now().Add(ttl)}
}

// NewDict creates DataType item with map value.
// Its DeathTime = (current time) + (provided TTL).
func NewDict(value map[interface{}]interface{}, ttl time.Duration) DataType {
	return DataType{value, ttl, time.Now().Add(ttl)}
}
