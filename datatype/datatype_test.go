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
	assert.Equal(t, dataType.ttl, duration)
	assert.Equal(t, dataType.DeathTime, expectedDeathTime)
}
