package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go-cache/datatype"
	"go-cache/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateString(t *testing.T) {
	Storage.Clear()

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateString).Methods(http.MethodPost)
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

func TestReadString(t *testing.T) {
	Storage.Clear()
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", ReadString).Methods(http.MethodGet)
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

func TestReadStringKeys(t *testing.T) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))

	router := mux.NewRouter()
	router.HandleFunc("/items/keys", ReadStringKeys).Methods(http.MethodGet)
	server := httptest.NewServer(router)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/items/keys", server.URL)
	request, err := http.NewRequest(http.MethodGet, usersUrl, nil)

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

func TestDeleteString(t *testing.T) {
	Storage.Clear()
	Storage.Set("name", datatype.NewString("Ivan", 2*time.Minute))

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", DeleteString).Methods(http.MethodDelete)
	server := httptest.NewServer(router)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/items/name", server.URL)
	request, err := http.NewRequest(http.MethodDelete, usersUrl, nil)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 204 {
		t.Errorf("HTTP Status expected: 204, got: %d", response.StatusCode)
	}
	assert.Equal(t, false, Storage.Contains("name"), "Key should not be present in storage")
}