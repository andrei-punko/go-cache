package datatype

import "time"

type DataType struct {
	Value     string `json:"value"`
	ttl       time.Duration
	DeathTime time.Time `json:"deathTime"`
}

func NewString(value string, duration time.Duration) DataType {
	return DataType{value, duration, time.Now().Add(duration)}
}
