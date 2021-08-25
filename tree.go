package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ruscalworld/nbt-cli/gotree"
	"github.com/urfave/cli/v2"
)

func PrintTree(_ *cli.Context) error {
	if len(CurrentData) == 0 {
		log.Println("This file does not contain any data.")
		return nil
	}

	root := gotree.New(fmt.Sprintf("• %s", filepath.Base(InputFilePath)))
	processNode(&root, CurrentData)
	root.SortItems()
	fmt.Println(root.Print())
	return nil
}

func processNode(parent *gotree.Tree, data map[string]interface{}) {
	for key, value := range data {
		if mapValue, ok := value.(map[string]interface{}); ok {
			processMap(parent, key, mapValue)
		} else if listValue, ok := value.([]interface{}); ok {
			processMap(parent, key, ArrayToMap(listValue))
		} else if intArrayValue, ok := value.([]int32); ok {
			processMap(parent, key, IntArrayToMap(intArrayValue))
		} else if longArrayValue, ok := value.([]int64); ok {
			processMap(parent, key, LongArrayToMap(longArrayValue))
		} else {
			(*parent).Add(fmt.Sprintf("%s: %s", key, ToString(value)))
		}
	}
}

func processMap(parent *gotree.Tree, key string, data map[string]interface{}) {
	child := (*parent).Add(key)
	processNode(&child, data)
	child.SortItems()
}
