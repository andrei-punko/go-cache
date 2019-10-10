package datastore

import (
	"github.com/umpc/go-sortedmap"
	"go-cache/datatype"
	"sync"
)

type DataStore struct {
	sync.RWMutex // ← this mutex protect cache below
	cache        sortedmap.SortedMap
}

func New() *DataStore {
	return &DataStore{
		cache: buildSortedMap(),
	}
}

func compareDataTypesByDeathTime(item1, item2 interface{}) bool {
	dt1 := item1.(datatype.DataType)
	dt2 := item2.(datatype.DataType)
	return dt1.DeathTime.Before(dt2.DeathTime)
}

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

func (ds *DataStore) Set(key string, value datatype.DataType) {
	ds.Lock()
	defer ds.Unlock()
	ds.set(key, value)
}

func (ds *DataStore) Get(key string) (interface{}, bool) {
	ds.RLock()
	defer ds.RUnlock()
	return ds.get(key)
}

func (ds *DataStore) GetKeys() []interface{} {
	ds.RLock()
	defer ds.RUnlock()
	return ds.getKeys()
}

func (ds *DataStore) Delete(key interface{}) bool {
	ds.Lock()
	defer ds.Unlock()
	return ds.delete(key)
}

func (ds *DataStore) BatchDelete(keys []interface{}) []bool {
	ds.Lock()
	defer ds.Unlock()
	return ds.batchDelete(keys)
}

func (ds *DataStore) Contains(key string) bool {
	ds.RLock()
	defer ds.RUnlock()
	return ds.contains(key)
}

func (ds *DataStore) Count() int {
	ds.RLock()
	defer ds.RUnlock()
	return ds.count()
}

func (ds *DataStore) Clear() {
	ds.Lock()
	defer ds.Unlock()
	ds.clear()
}

func (ds *DataStore) clear() {
	ds.cache = buildSortedMap()
}
