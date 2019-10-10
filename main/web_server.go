package main

import (
	"github.com/carlescere/scheduler"
	"github.com/gorilla/mux"
	json "github.com/json-iterator/go"
	"go-cache/datastore"
	"go-cache/datatype"
	"net/http"
	"os"
	"time"
)

var Storage = datastore.New()

// TODO: add load tests
func main() {
	scheduler.Every(10).Seconds().Run(cleanupExpiredItems)

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/keys", ReadKeys).Methods(http.MethodGet)
	router.HandleFunc("/items/keys", Clear).Methods(http.MethodDelete)
	router.HandleFunc("/items/{key}", ReadItem).Methods(http.MethodGet)
	router.HandleFunc("/items/{key}", DeleteItem).Methods(http.MethodDelete)

	http.ListenAndServe(":" + determinePort(), router)
}

func determinePort() string {
	if len(os.Args) == 1 {
		return "8000"
	}
	return os.Args[1]
}

// cleanupExpiredItems removes expired items from storage.
func cleanupExpiredItems() {
	// TODO: replace with more effective cleanup method
	keys := Storage.GetKeys()
	mostRightIndex := -1
	for index, key := range keys {
		value, _ := Storage.Get(key.(string))
		dataTypeItem := value.(datatype.DataType)
		if dataTypeItem.DeathTime.Before(time.Now()) {
			mostRightIndex = index
		} else {
			break
		}
	}
	if mostRightIndex != -1 {
		Storage.BatchDelete(keys[:mostRightIndex+1])
	}
}

// CreateItem creates item and saves it to storage.
func CreateItem(writer http.ResponseWriter, request *http.Request) {
	var value datatype.DataType
	err := json.NewDecoder(request.Body).Decode(&value)
	if err != nil {
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}
	value.DeathTime = time.Now().Add(value.Ttl)

	vars := mux.Vars(request)
	key := vars["key"]
	Storage.Set(key, value)
	resultJson, err := json.Marshal(value)
	if err != nil {
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}

	populateResponseWriter(writer, http.StatusCreated)
	writer.Write(resultJson)
}

// ReadItem reads item from storage and returns it.
func ReadItem(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	value, ok := Storage.Get(key)
	if !ok {
		populateResponseWriter(writer, http.StatusNotFound)
		return
	}

	resultJson, err := json.Marshal(value)
	if err != nil {
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}

	populateResponseWriter(writer, http.StatusOK)
	writer.Write(resultJson)
}

// ReadKeys reads and returns all keys saved in storage.
func ReadKeys(writer http.ResponseWriter, request *http.Request) {
	keys := Storage.GetKeys()
	resultJson, err := json.Marshal(keys)
	if err != nil {
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}

	populateResponseWriter(writer, http.StatusOK)
	writer.Write(resultJson)
}

// DeleteItem deletes specified item from storage.
func DeleteItem(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	if ok := Storage.Delete(key); !ok {
		populateResponseWriter(writer, http.StatusNotFound)
		return
	}

	populateResponseWriter(writer, http.StatusNoContent)
}

// Clear removes all items from storage.
func Clear(writer http.ResponseWriter, request *http.Request) {
	Storage.Clear()
	populateResponseWriter(writer, http.StatusNoContent)
}

// populateResponseWriter populates response header and status code.
func populateResponseWriter(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
}
