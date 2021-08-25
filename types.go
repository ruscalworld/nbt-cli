package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

func ToString(value interface{}) string {
	switch value.(type) {
	case float32, float64:
		return color.Blue.Text(fmt.Sprintf("%f", value))
	case uint8, int16, int32, int64:
		return color.Cyan.Text(fmt.Sprintf("%d", value))
	case string:
		if str, ok := value.(string); ok && !strings.Contains(str, "\"") {
			return color.Green.Text(fmt.Sprintf("\"%s\"", str))
		} else {
			return color.LightGreen.Text(fmt.Sprintf("'%s'", str))
		}
	default:
		return color.LightRed.Text(fmt.Sprintf("%s", value))
	}
}

func GetTypeName(value interface{}) string {
	switch value.(type) {
	case uint8:
		return "Byte"
	case int16:
		return "Short"
	case int32:
		return "Int"
	case int64:
		return "Long"
	case float32:
		return "Float"
	case float64:
		return "Double"
	case []byte:
		return "Byte Array"
	case string:
		return "String"
	case []interface{}:
		return "List"
	case map[string]interface{}:
		return "Compound"
	case []int32:
		return "Array of Int"
	case []int64:
		return "Array of Long"
	default:
		return "Unknown type"
	}
}

func ArrayToMap(array []interface{}) map[string]interface{} {
	if len(array) > ArrayItemLimit {
		omitted := len(array) - ArrayItemLimit
		array = array[:ArrayItemLimit]
		array = append(array, Tip{fmt.Sprintf("... and %d more items", omitted)})
	}

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
