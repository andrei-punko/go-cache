package main

import (
	"github.com/carlescere/scheduler"
	"github.com/gorilla/mux"
	json "github.com/json-iterator/go"
	"go-cache/datastore"
	"go-cache/datatype"
	"net/http"
	"time"
)

var Storage = datastore.New()

// TODO: add load tests
func main() {
	// TODO: populate items Storage externally, not in code
	Storage.Set("name", datatype.NewString("Roman", 1*time.Minute))
	Storage.Set("age", datatype.NewString("35", 5*time.Minute))
	Storage.Set("weight", datatype.NewString("82.5kg", 2*time.Minute))
	Storage.Set("car", datatype.NewString("Renault", 3*time.Minute))
	Storage.Set("phones", datatype.NewList([]interface{}{"Xiaomi", "Apple"}, 10*time.Minute))
	Storage.Set("cards", datatype.NewDict(map[interface{}]interface{}{2: "Visa", 3: "Maestro"}, 11*time.Minute))

	scheduler.Every(10).Seconds().Run(cleanupExpiredItems)

	router := mux.NewRouter()
	router.HandleFunc("/items/{key}", CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/items/keys", ReadKeys).Methods(http.MethodGet)
	router.HandleFunc("/items/{key}", ReadItem).Methods(http.MethodGet)
	router.HandleFunc("/items/{key}", DeleteItem).Methods(http.MethodDelete)

	http.ListenAndServe(":8000", router)
}

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

func DeleteItem(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	if ok := Storage.Delete(key); !ok {
		populateResponseWriter(writer, http.StatusNotFound)
		return
	}

	populateResponseWriter(writer, http.StatusNoContent)
}

func populateResponseWriter(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
}
