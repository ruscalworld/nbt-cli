package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/urfave/cli/v2"
)

func SetValue(context *cli.Context) error {
	key, value := context.String("key"), context.String("value")
	path := strings.Split(key, ".")

	target, err := walk(CurrentData, path)
	if err != nil {
		return err
	}

	target[path[len(path)-1]] = value
	return SaveData()
}

func walk(data map[string]interface{}, path []string) (map[string]interface{}, error) {
	for len(path) > 1 {
		value := data[path[0]]
		if mapValue, ok := value.(map[string]interface{}); ok {
			return walk(mapValue, path[1:])
		} else if value == nil {
			data[path[0]] = make(map[string]interface{})
			return walk(data, path)
		} else {
			return nil, errors.New(fmt.Sprintf("%s was a %s, but map was expected", path[0], reflect.TypeOf(value)))
		}
	}

	return data, nil
}
