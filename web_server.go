package main

import (
	"encoding/json"
	"github.com/carlescere/scheduler"
	"github.com/gorilla/mux"
	"go-sandbox/cache/datastore"
	"go-sandbox/cache/datatype"
	"net/http"
	"time"
)

var storage = datastore.New()

// TODO: add tests and load tests
func main() {
	// TODO: populate items storage externally, not in code
	storage.Set("name", datatype.New("Roman", 1*time.Minute))
	storage.Set("age", datatype.New("35", 5*time.Minute))
	storage.Set("weight", datatype.New("82.5kg", 2*time.Minute))
	storage.Set("car", datatype.New("Renault", 3*time.Minute))

	scheduler.Every(10).Seconds().Run(cleanupExpiredItems)

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateString).Methods(http.MethodPost)
	router.HandleFunc("/items/keys", ReadStringKeys).Methods(http.MethodGet)
	router.HandleFunc("/items/{key}", ReadString).Methods(http.MethodGet)
	router.HandleFunc("/items/{key}", DeleteString).Methods(http.MethodDelete)

	http.ListenAndServe(":8000", router)
}

func cleanupExpiredItems() {
	// TODO: replace with more effective cleanup method
	keys := storage.GetKeys()
	mostRightIndex := -1
	for index, key := range keys {
		value, _ := storage.Get(key.(string))
		dataTypeItem := value.(datatype.DataType)
		if dataTypeItem.DeathTime.Before(time.Now()) {
			mostRightIndex = index
		} else {
			break
		}
	}
	if mostRightIndex != -1 {
		storage.BatchDelete(keys[:mostRightIndex+1])
	}
}

func CreateString(writer http.ResponseWriter, request *http.Request) {
	var value datatype.DataType
	err := json.NewDecoder(request.Body).Decode(&value)
	if err != nil {
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(request)
	key := vars["key"]
	storage.Set(key, value)
	resultJson, err := json.Marshal(value)
	if err != nil {
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}

	populateResponseWriter(writer, http.StatusCreated)
	writer.Write(resultJson)
}

func ReadString(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	value, ok := storage.Get(key)
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

func ReadStringKeys(writer http.ResponseWriter, request *http.Request) {
	keys := storage.GetKeys()
	resultJson, err := json.Marshal(keys)
	if err != nil {
		populateResponseWriter(writer, http.StatusInternalServerError)
		return
	}

	populateResponseWriter(writer, http.StatusOK)
	writer.Write(resultJson)
}

func DeleteString(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	if ok := storage.Delete(key); !ok {
		populateResponseWriter(writer, http.StatusNotFound)
		return
	}

	populateResponseWriter(writer, http.StatusNoContent)
}

func populateResponseWriter(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
}
