package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceListToStringList(t *testing.T) {
	array1 := []interface{}{"str 1", "str 2"}

	array2 := InterfaceListToStringList(array1)
	expectedArray := []string{"str 1", "str 2"}
	assert.Equal(t, expectedArray, array2)
}

func TestStringListToInterfaceList(t *testing.T) {
	array1 := []string{"str 1", "str 2"}

	array2 := StringListToInterfaceList(array1)
	expectedArray := []interface{}{"str 1", "str 2"}
	assert.Equal(t, expectedArray, array2)
}
