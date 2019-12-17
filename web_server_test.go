package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrei-punko/go-cache/datatype"
	"github.com/andrei-punko/go-cache/util"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateItem(t *testing.T) {
	Storage.Clear()

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	server := httptest.NewServer(router)
	defer server.Close()
	itemsUrl := fmt.Sprintf("%s/items/name", server.URL)
	itemJson := `{"value": "Ioann", "ttl": 60000000000}`
	request, err := http.NewRequest(http.MethodPost, itemsUrl, strings.NewReader(itemJson))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", response.StatusCode)
	}
	var decodedObject datatype.DataType
	err = json.NewDecoder(response.Body).Decode(&decodedObject)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Ioann", decodedObject.Value, "Wrong value in decoded object")
	assert.Equal(t, time.Minute, decodedObject.Ttl, "Wrong ttl in decoded object")
	assert.LessOrEqual(t, (time.Now().Add(decodedObject.Ttl).Sub(decodedObject.DeathTime)).Seconds(), 0.1)

	value, ok := Storage.Get("name")
	dataTypeItem := value.(datatype.DataType)
	assert.Equal(t, true, ok, "Item should be present in storage")
	assert.Equal(t, "Ioann", dataTypeItem.Value, "Wrong value")
	assert.Equal(t, time.Minute, dataTypeItem.Ttl, "Wrong ttl")
	assert.LessOrEqual(t, (time.Now().Add(dataTypeItem.Ttl).Sub(dataTypeItem.DeathTime)).Seconds(), 0.1)
}

func TestReadItem(t *testing.T) {
	Storage.Clear()
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", ReadItem).Methods(http.MethodGet)
	server := httptest.NewServer(router)
	defer server.Close()
	itemsUrl := fmt.Sprintf("%s/items/weight", server.URL)
	request, err := http.NewRequest(http.MethodGet, itemsUrl, nil)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", response.StatusCode)
	}
	var decodedObject datatype.DataType
	err = json.NewDecoder(response.Body).Decode(&decodedObject)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "82.5kg", decodedObject.Value, "Wrong value in decoded object")
	assert.Equal(t, 2*time.Minute, decodedObject.Ttl, "Wrong ttl in decoded object")
	assert.LessOrEqual(t, (time.Now().Add(decodedObject.Ttl).Sub(decodedObject.DeathTime)).Seconds(), 0.1)
}

func TestReadKeys(t *testing.T) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	router := mux.NewRouter()
	router.HandleFunc("/items/keys", ReadKeys).Methods(http.MethodGet)
	server := httptest.NewServer(router)
	defer server.Close()
	itemsUrl := fmt.Sprintf("%s/items/keys", server.URL)
	request, err := http.NewRequest(http.MethodGet, itemsUrl, nil)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", response.StatusCode)
	}
	var decodedObject []interface{}
	err = json.NewDecoder(response.Body).Decode(&decodedObject)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, len(decodedObject), "Array should contains 2 keys")
	assert.Equal(t, true, util.ContainsAll(decodedObject, []interface{}{"name", "weight"}))
}

func TestDeleteItem(t *testing.T) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", DeleteItem).Methods(http.MethodDelete)
	server := httptest.NewServer(router)
	defer server.Close()
	itemsUrl := fmt.Sprintf("%s/items/name", server.URL)
	request, err := http.NewRequest(http.MethodDelete, itemsUrl, nil)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 204 {
		t.Errorf("HTTP Status expected: 204, got: %d", response.StatusCode)
	}
	assert.Equal(t, false, Storage.Contains("name"), "Key should not be present in storage")
}

func TestClear(t *testing.T) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	router := mux.NewRouter()
	router.HandleFunc("/items/keys", Clear).Methods(http.MethodDelete)
	server := httptest.NewServer(router)
	defer server.Close()
	itemsUrl := fmt.Sprintf("%s/items/keys", server.URL)
	request, err := http.NewRequest(http.MethodDelete, itemsUrl, nil)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 204 {
		t.Errorf("HTTP Status expected: 204, got: %d", response.StatusCode)
	}
	assert.Equal(t, 0, Storage.Count(), "Storage should be empty")
}

func Test_determineIndexForCleanup(t *testing.T) {
	Storage.Clear()
	Storage.Set("name1", datatype.NewString("Ivan", 1*time.Second))
	Storage.Set("name2", datatype.NewString("Ivan", 2*time.Second))
	Storage.Set("name3", datatype.NewString("Ivan", 3*time.Second))
	Storage.Set("name4", datatype.NewString("Ivan", 4*time.Second))
	Storage.Set("name5", datatype.NewString("Ivan", 5*time.Second))
	Storage.Set("name6", datatype.NewString("Ivan", 6*time.Second))
	Storage.Set("name7", datatype.NewString("Ivan", 7*time.Second))

	assert.Equal(t, -1, determineIndexForCleanup(Storage.GetKeys(), time.Now()))
	assert.Equal(t, 1, determineIndexForCleanup(Storage.GetKeys(), time.Now().Add(2500*time.Millisecond)))
	assert.Equal(t, 2, determineIndexForCleanup(Storage.GetKeys(), time.Now().Add(3500*time.Millisecond)))
}

func Test_isBefore(t *testing.T) {
	Storage.Clear()
	Storage.Set("name1", datatype.NewString("Ivan", 1*time.Second))
	Storage.Set("name2", datatype.NewString("Ivan", 2*time.Second))
	Storage.Set("name3", datatype.NewString("Ivan", 3*time.Second))
	Storage.Set("name4", datatype.NewString("Ivan", 4*time.Second))
	Storage.Set("name5", datatype.NewString("Ivan", 5*time.Second))

	assert.Equal(t, true, isBefore("name1", time.Now().Add(2500*time.Millisecond)))
	assert.Equal(t, true, isBefore("name2", time.Now().Add(2500*time.Millisecond)))
	assert.Equal(t, false, isBefore("name3", time.Now().Add(2500*time.Millisecond)))
	assert.Equal(t, false, isBefore("name4", time.Now().Add(2500*time.Millisecond)))
	assert.Equal(t, false, isBefore("name5", time.Now().Add(2500*time.Millisecond)))
}

func ExampleCreateItem() {
	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/{key}", ReadItem).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}

func ExampleReadItem() {
	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/{key}", ReadItem).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}

func ExampleReadKeys() {
	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/keys", ReadKeys).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}

func ExampleDeleteItem() {
	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/{key}", DeleteItem).Methods(http.MethodDelete)
	http.ListenAndServe(":8000", router)
}

func ExampleClear() {
	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/keys", Clear).Methods(http.MethodDelete)
	http.ListenAndServe(":8000", router)
}

func BenchmarkCreateItem(b *testing.B) {
	Storage.Clear()

	writer := datatype.NewStubResponseWriter()
	itemJson := `{"value": "Ioann", "ttl": 60000000000}`
	reader := strings.NewReader(itemJson)
	request, _ := http.NewRequest(http.MethodPost, "someUrl", reader)

	for n := 0; n < b.N; n++ {
		CreateItem(writer, request)
		// Need to reset reader because it could be used once
		reader.Seek(0, io.SeekStart)
	}
}

func BenchmarkReadItem(b *testing.B) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	writer := datatype.NewStubResponseWriter()
	request, _ := http.NewRequest(http.MethodGet, "someUrl", nil)
	request = mux.SetURLVars(request, map[string]string{"key": "weight"})

	for n := 0; n < b.N; n++ {
		ReadItem(writer, request)
	}
}

func BenchmarkReadKeys(b *testing.B) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	writer := datatype.NewStubResponseWriter()
	request, _ := http.NewRequest(http.MethodGet, "someUrl", nil)

	for n := 0; n < b.N; n++ {
		ReadKeys(writer, request)
	}
}

func BenchmarkDeleteItem(b *testing.B) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	writer := datatype.NewStubResponseWriter()
	request, _ := http.NewRequest(http.MethodDelete, "someUrl", nil)
	request = mux.SetURLVars(request, map[string]string{"key": "name"})

	for n := 0; n < b.N; n++ {
		DeleteItem(writer, request)
	}
}

func BenchmarkClear(b *testing.B) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	writer := datatype.NewStubResponseWriter()
	request, _ := http.NewRequest(http.MethodDelete, "someUrl", nil)

	for n := 0; n < b.N; n++ {
		Clear(writer, request)
	}
}
