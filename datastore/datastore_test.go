package datastore

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-cache/datatype"
	"go-cache/util"
	"testing"
	"time"
)

func TestNewDataStore(t *testing.T) {
	dataStore := NewDataStore()
	assert.Equal(t, dataStore.cache.Len(), 0, "map should be empty")
}

func TestDataStore_set(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)

	dataStore.set(key, value)
	actualValue, _ := dataStore.cache.Get(key)
	assert.Equal(t, value, actualValue)
}

func TestDataStore_get(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	actualValue, ok := dataStore.get(key)
	assert.Equal(t, value, actualValue)
	assert.Equal(t, true, ok)
}

func TestDataStore_getKeys(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	keys := dataStore.getKeys()
	assert.Equal(t, len(keys), 2)
	assert.Contains(t, keys, "key 1")
	assert.Contains(t, keys, "key 2")
}

func TestDataStore_delete(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.delete(key), true, "existing key")
	assert.Equal(t, dataStore.cache.Len(), 0, "map should be empty")
	assert.Equal(t, dataStore.delete(key), false, "absent key")
	assert.Equal(t, dataStore.delete("another key"), false, "absent key 2")
}

func TestDataStore_batchDelete(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))
	dataStore.cache.Insert("key 3", datatype.NewString("value 3", time.Minute))
	keys := util.StringListToInterfaceList([]string{"key 1", "key N", "key 2", "key Z"})

	results := dataStore.batchDelete(keys)
	assert.Equal(t, 4, len(results), "Two items in result array expected")
	assert.Equal(t, []bool{true, false, true, false}, results)
	assert.Equal(t, 1, dataStore.cache.Len(), "One item should remain")
	assert.Equal(t, true, dataStore.cache.Has("key 3"))
}

func TestDataStore_BatchDelete(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))
	dataStore.cache.Insert("key 3", datatype.NewString("value 3", time.Minute))
	keys := util.StringListToInterfaceList([]string{"key 1", "key N", "key 2", "key Z"})

	results := dataStore.BatchDelete(keys)
	assert.Equal(t, 4, len(results), "Two items in result array expected")
	assert.Equal(t, []bool{true, false, true, false}, results)
	assert.Equal(t, 1, dataStore.cache.Len(), "One item should remain")
	assert.Equal(t, true, dataStore.cache.Has("key 3"))
}

func TestDataStore_contains(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.contains(key), true, "existing key")
	assert.Equal(t, dataStore.contains("another key"), false, "absent key")
}

func TestDataStore_count(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	assert.Equal(t, dataStore.count(), 2)
}

func TestDataStore_clear(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	dataStore.clear()
	assert.Equal(t, 0, dataStore.cache.Len(), "Storage should be empty")
}

func TestDataStore_Set(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)

	dataStore.Set(key, value)
	actualValue, _ := dataStore.cache.Get(key)
	assert.Equal(t, value, actualValue)
}

func TestDataStore_Get(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	res, ok := dataStore.Get(key)
	assert.Equal(t, value, res)
	assert.Equal(t, true, ok)

	res, ok = dataStore.Get("another key")
	if res != nil {
		t.Errorf("nil should be returned")
	}
	assert.Equal(t, false, ok)
}

func TestDataStore_GetKeys(t *testing.T) {
	dataStore := NewDataStore()
	keys := dataStore.GetKeys()
	assert.Equal(t, []interface{}{}, keys, "Should return empty array in case of empty storage")

	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	keys = dataStore.GetKeys()
	assert.Equal(t, len(keys), 2)
	assert.Contains(t, keys, "key 1")
	assert.Contains(t, keys, "key 2")
}

func TestDataStore_Delete(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.Delete(key), true, "existing key")
	assert.Equal(t, dataStore.cache.Len(), 0, "map should be empty")
	assert.Equal(t, dataStore.Delete(key), false, "absent key")
	assert.Equal(t, dataStore.Delete("another key"), false, "absent key 2")
}

func TestDataStore_Contains(t *testing.T) {
	dataStore := NewDataStore()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.Contains(key), true, "existing key")
	assert.Equal(t, dataStore.Contains("another key"), false, "absent key")
}

func TestDataStore_Count(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	assert.Equal(t, dataStore.Count(), 2)
}

func TestDataStore_Clear(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	dataStore.Clear()
	assert.Equal(t, 0, dataStore.cache.Len(), "Storage should be empty")
}

func TestDataStore_compareDataTypesByDeathTime(t *testing.T) {
	dt1 := datatype.NewString("value 1", time.Minute)
	dt2 := datatype.NewString("value 2", 2*time.Minute)
	assert.Equal(t, true, compareDataTypesByDeathTime(dt1, dt2))
}

func ExampleDataStore_Set() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("phones", datatype.NewList([]interface{}{"Xiaomi", "Samsung"}, 2*time.Minute))
	storage.Set("cards", datatype.NewDict(map[interface{}]interface{}{2: "Visa", 3: "Maestro"}, 4*time.Minute))
	fmt.Println(storage.GetKeys())
}

func ExampleDataStore_Get() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	fmt.Println(storage.Get("name"))
}

func ExampleDataStore_GetKeys() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	fmt.Println(storage.GetKeys())
}

func ExampleDataStore_Delete() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	storage.Delete("name")
	fmt.Println(storage.GetKeys())
}

func ExampleDataStore_BatchDelete() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	storage.Set("weight", datatype.NewString("80.5kg", time.Minute))
	storage.BatchDelete([]interface{}{"name", "age"})
	fmt.Println(storage.GetKeys())
}

func ExampleDataStore_Contains() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	fmt.Print(storage.Contains("name"))
	fmt.Print(storage.Contains("weight"))
}

func ExampleDataStore_Count() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	fmt.Println(storage.Count())
}

func ExampleDataStore_Clear() {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	storage.Clear()
}

func BenchmarkNewDataStore(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewDataStore()
	}
}

func BenchmarkDataStore_Set(b *testing.B) {
	storage := NewDataStore()

	for n := 0; n < b.N; n++ {
		storage.Set("name", datatype.NewString("Ivan", time.Minute))
	}
}

func BenchmarkDataStore_Get(b *testing.B) {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))

	for n := 0; n < b.N; n++ {
		storage.Get("name")
	}
}

func BenchmarkDataStore_GetKeys(b *testing.B) {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))

	for n := 0; n < b.N; n++ {
		storage.GetKeys()
	}
}

func BenchmarkDataStore_Contains(b *testing.B) {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))

	for n := 0; n < b.N; n++ {
		storage.Contains("name")
	}
}

func BenchmarkDataStore_Count(b *testing.B) {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))

	for n := 0; n < b.N; n++ {
		storage.Count()
	}
}

func BenchmarkDataStore_Delete(b *testing.B) {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))

	for n := 0; n < b.N; n++ {
		storage.Delete("name")
	}
}

func BenchmarkDataStore_BatchDelete(b *testing.B) {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))
	storage.Set("weight", datatype.NewString("80.5kg", time.Minute))

	for n := 0; n < b.N; n++ {
		storage.BatchDelete([]interface{}{"name", "age"})
	}
}

func BenchmarkDataStore_Clear(b *testing.B) {
	storage := NewDataStore()
	storage.Set("name", datatype.NewString("Ivan", time.Minute))
	storage.Set("age", datatype.NewString("27", time.Minute))

	for n := 0; n < b.N; n++ {
		storage.Clear()
	}
}
