package util

import "fmt"

func InterfaceListToStringList(list []interface{}) []string {
	res := make([]string, len(list))
	for i, v := range list {
		res[i] = fmt.Sprint(v)
	}
	return res
}

func StringListToInterfaceList(list []string) []interface{} {
	res := make([]interface{}, len(list))
	for i, v := range list {
		res[i] = v
	}
	return res
}

func Contains(arr []interface{}, item interface{}) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}

func ContainsAll(arr []interface{}, items []interface{}) bool {
	for _, item := range items {
		if !Contains(arr, item) {
			return false
		}
	}
	return true
}
