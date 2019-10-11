package datatype

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewString(t *testing.T) {
	value := "value 2"
	duration := time.Minute
	expectedDeathTime := time.Now().Add(duration)

	dataType := NewString(value, duration)
	assert.Equal(t, dataType.Value, value)
	assert.Equal(t, dataType.Ttl, duration)
	assert.Equal(t, dataType.DeathTime, expectedDeathTime)
}

func TestNewList(t *testing.T) {
	value := []interface{}{2, 5, 9}
	duration := time.Minute
	expectedDeathTime := time.Now().Add(duration)

	dataType := NewList(value, duration)
	assert.Equal(t, dataType.Value, value)
	assert.Equal(t, dataType.Ttl, duration)
	assert.Equal(t, dataType.DeathTime, expectedDeathTime)
}

func TestNewDict(t *testing.T) {
	value := map[interface{}]interface{}{2: "two", 5: "five"}
	duration := time.Minute
	expectedDeathTime := time.Now().Add(duration)

	dataType := NewDict(value, duration)
	assert.Equal(t, dataType.Value, value)
	assert.Equal(t, dataType.Ttl, duration)
	assert.Equal(t, dataType.DeathTime, expectedDeathTime)
}

func ExampleNewString() {
	value := "value 2"
	duration := time.Minute
	NewString(value, duration)
}

func ExampleNewList() {
	value := []interface{}{2, 5, 9}
	duration := time.Minute
	NewList(value, duration)
}

func ExampleNewDict() {
	value := map[interface{}]interface{}{2: "two", 5: "five"}
	duration := time.Minute
	NewDict(value, duration)
}
