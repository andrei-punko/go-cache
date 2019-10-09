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

func TestContains(t *testing.T) {
	arr := []interface{}{"one", "five"}

	assert.Equal(t, true, Contains(arr, "one"), "Should contain first item")
	assert.Equal(t, true, Contains(arr, "five"), "Should contain second item")
	assert.Equal(t, false, Contains(arr, "six"), "Shouldn't contain another item")
}

func TestContainsAll(t *testing.T) {
	arr := []interface{}{"one", "five"}
	items := []interface{}{"one", "five"}
	items2 := []interface{}{"one", "six"}

	assert.Equal(t, true, ContainsAll(arr, items), "Should contain all items from set")
	assert.Equal(t, false, ContainsAll(arr, items2), "Should not contain all items from another set")
}

func TestRandString(t *testing.T) {
	assert.NotEqual(t, RandString(10), RandString(10), "Generated strings should not be same")
}
