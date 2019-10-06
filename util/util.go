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
