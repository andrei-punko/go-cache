package datatype

import "time"

type DataType struct {
	Value     interface{}   `json:"value"`
	Ttl       time.Duration `json:"ttl"`
	DeathTime time.Time     `json:"deathTime"`
}

func NewString(value string, ttl time.Duration) DataType {
	return DataType{value, ttl, time.Now().Add(ttl)}
}

func NewList(value []interface{}, ttl time.Duration) DataType {
	return DataType{value, ttl, time.Now().Add(ttl)}
}

func NewDict(value map[interface{}]interface{}, ttl time.Duration) DataType {
	return DataType{value, ttl, time.Now().Add(ttl)}
}
