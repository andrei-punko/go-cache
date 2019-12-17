package main

import (
	"github.com/carlescere/scheduler"
	"github.com/gorilla/mux"
	json "github.com/json-iterator/go"
	"github.com/andrei-punko/go-cache/datastore"
	"github.com/andrei-punko/go-cache/datatype"
	"log"
	"net/http"
	"os"
	"time"
)

var Storage = datastore.NewDataStore()

func main() {
	port := extractPortFromCmdParams()
	log.Printf("Starting web-cache on port %s ...", port)
	scheduler.Every(10).Seconds().Run(cleanupExpiredItems)

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/keys", ReadKeys).Methods(http.MethodGet)
	router.HandleFunc("/items/keys", Clear).Methods(http.MethodDelete)
	router.HandleFunc("/items/{key}", ReadItem).Methods(http.MethodGet)
	router.HandleFunc("/items/{key}", DeleteItem).Methods(http.MethodDelete)

	http.ListenAndServe(":"+port, router)
}

func extractPortFromCmdParams() string {
	if len(os.Args) == 1 {
		return "8000"
	}
	return os.Args[1]
}

// cleanupExpiredItems removes expired items from storage.
func cleanupExpiredItems() {
	keys := Storage.GetKeys()
	indexForCleanup := determineIndexForCleanup(keys, time.Now())
	if indexForCleanup != -1 {
		Storage.BatchDelete(keys[:indexForCleanup+1])
	}
}

// determineIndexForCleanup used binary search to determine rightmost index of expired items.
func determineIndexForCleanup(keys []interface{}, time time.Time) int {
	leftIndex := -1
	rightIndex := len(keys) - 1
	for rightIndex-leftIndex > 1 {
		index := (leftIndex + rightIndex) / 2
		if isBefore(keys[index], time) {
			leftIndex = index
		} else {
			rightIndex = index
		}
	}

	return leftIndex
}

func isBefore(key interface{}, time time.Time) bool {
	value, _ := Storage.Get(key.(string))
	dataTypeItem := value.(datatype.DataType)
	return dataTypeItem.DeathTime.Before(time)
}

// CreateItem creates item and saves it to storage.
func CreateItem(writer http.ResponseWriter, request *http.Request) {
	var value datatype.DataType
	err := json.NewDecoder(request.Body).Decode(&value)
	if err != nil {
		log.Println("Error during json decoding")
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}
	value.DeathTime = time.Now().Add(value.Ttl)

	vars := mux.Vars(request)
	key := vars["key"]
	Storage.Set(key, value)
	resultJson, err := json.Marshal(value)
	if err != nil {
		log.Println("Error during json encoding")
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
		log.Println("Error during json encoding")
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
		log.Println("Error during json encoding")
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
