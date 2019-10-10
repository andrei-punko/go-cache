package util

import (
	"fmt"
	"math/rand"
	"time"
)

// InterfaceListToStringList converts list of interfaces to list of strings.
func InterfaceListToStringList(list []interface{}) []string {
	res := make([]string, len(list))
	for i, v := range list {
		res[i] = fmt.Sprint(v)
	}
	return res
}

// StringListToInterfaceList converts list of strings to list of interfaces.
func StringListToInterfaceList(list []string) []interface{} {
	res := make([]interface{}, len(list))
	for i, v := range list {
		res[i] = v
	}
	return res
}

// Contains returns flag is array contains provided item.
func Contains(arr []interface{}, item interface{}) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}

// ContainsAll returns flag is array contains each of provided items.
func ContainsAll(arr []interface{}, items []interface{}) bool {
	for _, item := range items {
		if !Contains(arr, item) {
			return false
		}
	}
	return true
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generate random string with definite length.
// According to https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
