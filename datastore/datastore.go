package datastore

import (
	"github.com/umpc/go-sortedmap"
	"go-cache/datatype"
	"sync"
)

// DataStore contains map and mutex to protect it.
type DataStore struct {
	sync.RWMutex
	cache sortedmap.SortedMap
}

// New creates and initializes a new DataStore structure and then returns a reference to it.
// DataStore is concurrency-safe.
func New() *DataStore {
	return &DataStore{cache: buildSortedMap()}
}

// compareDataTypesByDeathTime compares DataType items by DeathTime.
func compareDataTypesByDeathTime(item1, item2 interface{}) bool {
	dt1 := item1.(datatype.DataType)
	dt2 := item2.(datatype.DataType)
	return dt1.DeathTime.Before(dt2.DeathTime)
}

// buildSortedMap creates SortedMap with size 10 and comparison function which compares items by DeathTime.
func buildSortedMap() sortedmap.SortedMap {
	return *sortedmap.New(10, compareDataTypesByDeathTime)
}

func (ds *DataStore) set(key string, value datatype.DataType) {
	ds.cache.Replace(key, value)
}

func (ds *DataStore) get(key string) (interface{}, bool) {
	return ds.cache.Get(key)
}

func (ds *DataStore) getKeys() []interface{} {
	return ds.cache.Keys()
}

func (ds *DataStore) delete(key interface{}) bool {
	return ds.cache.Delete(key)
}

func (ds *DataStore) batchDelete(keys []interface{}) []bool {
	return ds.cache.BatchDelete(keys)
}

func (ds *DataStore) contains(key string) bool {
	_, ok := ds.cache.Get(key)
	return ok
}

func (ds *DataStore) count() int {
	return ds.cache.Len()
}

// Set adds provided key-value pair to the collection.
func (ds *DataStore) Set(key string, value datatype.DataType) {
	ds.Lock()
	defer ds.Unlock()
	ds.set(key, value)
}

// Get returns value for provided key stored in the collection.
func (ds *DataStore) Get(key string) (interface{}, bool) {
	ds.RLock()
	defer ds.RUnlock()
	return ds.get(key)
}

// GetKeys returns all keys stored in the collection.
func (ds *DataStore) GetKeys() []interface{} {
	ds.RLock()
	defer ds.RUnlock()
	return ds.getKeys()
}

// Delete deletes provided key from the collection.
func (ds *DataStore) Delete(key interface{}) bool {
	ds.Lock()
	defer ds.Unlock()
	return ds.delete(key)
}

// BatchDelete deletes provided keys from the collection.
func (ds *DataStore) BatchDelete(keys []interface{}) []bool {
	ds.Lock()
	defer ds.Unlock()
	return ds.batchDelete(keys)
}

// Contains returns flag is this key present in the collection.
func (ds *DataStore) Contains(key string) bool {
	ds.RLock()
	defer ds.RUnlock()
	return ds.contains(key)
}

// Count returns amount of items in the collection.
func (ds *DataStore) Count() int {
	ds.RLock()
	defer ds.RUnlock()
	return ds.count()
}

// Clear removes all items from the collection.
func (ds *DataStore) Clear() {
	ds.Lock()
	defer ds.Unlock()
	ds.clear()
}

func (ds *DataStore) clear() {
	ds.cache = buildSortedMap()
}
