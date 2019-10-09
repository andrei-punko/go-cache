package datastore

import (
	"github.com/stretchr/testify/assert"
	"go-cache/datatype"
	"go-cache/util"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	dataStore := New()
	assert.Equal(t, dataStore.cache.Len(), 0, "map should be empty")
}

func TestDataStore_set(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)

	dataStore.set(key, value)
	actualValue, _ := dataStore.cache.Get(key)
	assert.Equal(t, value, actualValue)
}

func TestDataStore_get(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	actualValue, ok := dataStore.get(key)
	assert.Equal(t, value, actualValue)
	assert.Equal(t, true, ok)
}

func TestDataStore_getKeys(t *testing.T) {
	dataStore := New()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	keys := dataStore.getKeys()
	assert.Equal(t, len(keys), 2)
	assert.Contains(t, keys, "key 1")
	assert.Contains(t, keys, "key 2")
}

func TestDataStore_delete(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.delete(key), true, "existing key")
	assert.Equal(t, dataStore.cache.Len(), 0, "map should be empty")
	assert.Equal(t, dataStore.delete(key), false, "absent key")
	assert.Equal(t, dataStore.delete("another key"), false, "absent key 2")
}

func TestDataStore_batchDelete(t *testing.T) {
	dataStore := New()
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
	dataStore := New()
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
	dataStore := New()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.contains(key), true, "existing key")
	assert.Equal(t, dataStore.contains("another key"), false, "absent key")
}

func TestDataStore_count(t *testing.T) {
	dataStore := New()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	assert.Equal(t, dataStore.count(), 2)
}

func TestDataStore_clear(t *testing.T) {
	dataStore := New()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	dataStore.clear()
	assert.Equal(t, 0, dataStore.cache.Len(), "Storage should be empty")
}

func TestDataStore_Set(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)

	dataStore.Set(key, value)
	actualValue, _ := dataStore.cache.Get(key)
	assert.Equal(t, value, actualValue)
}

func TestDataStore_Get(t *testing.T) {
	dataStore := New()
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
	dataStore := New()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	keys := dataStore.GetKeys()
	assert.Equal(t, len(keys), 2)
	assert.Contains(t, keys, "key 1")
	assert.Contains(t, keys, "key 2")
}

func TestDataStore_Delete(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.Delete(key), true, "existing key")
	assert.Equal(t, dataStore.cache.Len(), 0, "map should be empty")
	assert.Equal(t, dataStore.Delete(key), false, "absent key")
	assert.Equal(t, dataStore.Delete("another key"), false, "absent key 2")
}

func TestDataStore_Contains(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.NewString("Some value", time.Minute)
	dataStore.cache.Insert(key, value)

	assert.Equal(t, dataStore.Contains(key), true, "existing key")
	assert.Equal(t, dataStore.Contains("another key"), false, "absent key")
}

func TestDataStore_Count(t *testing.T) {
	dataStore := New()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	assert.Equal(t, dataStore.Count(), 2)
}

func TestDataStore_Clear(t *testing.T) {
	dataStore := New()
	dataStore.cache.Insert("key 1", datatype.NewString("value 1", time.Minute))
	dataStore.cache.Insert("key 2", datatype.NewString("value 2", time.Minute))

	dataStore.Clear()
	assert.Equal(t, 0, dataStore.cache.Len(), "Storage should be empty")
}

func TestDataStore_compareDataTypeItems(t *testing.T) {
	dt1 := datatype.NewString("value 1", time.Minute)
	dt2 := datatype.NewString("value 2", 2*time.Minute)
	assert.Equal(t, true, compareDataTypeItems(dt1, dt2))
}
