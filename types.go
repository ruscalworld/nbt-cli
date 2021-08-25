package main

import (
	"fmt"
	"strconv"
)

func ToString(value interface{}) string {
	switch value.(type) {
	case float32, float64:
		return fmt.Sprintf("%f", value)
	case uint8, int16, int32, int64:
		return fmt.Sprintf("%d", value)
	default:
		return fmt.Sprintf("%s", value)
	}
}

func ArrayToMap(array []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, v := range array {
		result[strconv.Itoa(i)] = v
	}
	return result
}

func IntArrayToMap(array []int32) map[string]interface{} {
	result := make([]interface{}, 0)
	for _, v := range array {
		result = append(result, v)
	}
	return ArrayToMap(result)
}

func LongArrayToMap(array []int64) map[string]interface{} {
	result := make([]interface{}, 0)
	for _, v := range array {
		result = append(result, v)
	}
	return ArrayToMap(result)
}
