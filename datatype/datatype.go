package datatype

import "time"

type DataType struct {
	Value     interface{} `json:"value"`
	ttl       time.Duration
	DeathTime time.Time `json:"deathTime"`
}

func NewString(value string, duration time.Duration) DataType {
	return DataType{value, duration, time.Now().Add(duration)}
}

func NewList(value []interface{}, duration time.Duration) DataType {
	return DataType{value, duration, time.Now().Add(duration)}
}

func NewDict(value map[interface{}]interface{}, duration time.Duration) DataType {
	return DataType{value, duration, time.Now().Add(duration)}
}
