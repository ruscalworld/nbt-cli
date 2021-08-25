package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"path/filepath"
	"time"

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
			processMap(parent, key, "", mapValue)
		} else if listValue, ok := value.([]interface{}); ok {
			processMap(parent, key, "", ArrayToMap(listValue))
		} else if intArrayValue, ok := value.([]int32); ok {
			comment := ""

			if len(intArrayValue) == 4 {
				parsedUUID, err := uuid.FromBytes([]byte{
					byte(intArrayValue[0] >> 24), byte(intArrayValue[0] >> 16), byte(intArrayValue[0] >> 8), byte(intArrayValue[0]),
					byte(intArrayValue[1] >> 24), byte(intArrayValue[1] >> 16), byte(intArrayValue[1] >> 8), byte(intArrayValue[1]),
					byte(intArrayValue[2] >> 24), byte(intArrayValue[2] >> 16), byte(intArrayValue[2] >> 8), byte(intArrayValue[2]),
					byte(intArrayValue[3] >> 24), byte(intArrayValue[3] >> 16), byte(intArrayValue[3] >> 8), byte(intArrayValue[3]),
				})

				if err == nil {
					comment = Comment(fmt.Sprintf("UUID: %s", parsedUUID.String()))
				}
			}

			processMap(parent, key, comment, IntArrayToMap(intArrayValue))
		} else if longArrayValue, ok := value.([]int64); ok {
			processMap(parent, key, "", LongArrayToMap(longArrayValue))
		} else if tip, ok := value.(Tip); ok {
			(*parent).Add(tip.Text)
		} else {
			comment := ""
			if long, ok := value.(int64); ok {
				if long > 1000000000000 {
					t := time.Unix(long/1000, 0)
					comment = Comment(fmt.Sprintf("Time: %s", t))
				}
			}

			(*parent).Add(fmt.Sprintf("%s: %s %s", key, ToString(value), comment))
		}
	}
}

func processMap(parent *gotree.Tree, key, description string, data map[string]interface{}) {
	child := (*parent).Add(fmt.Sprintf("%s %s", key, description))
	processNode(&child, data)
	child.SortItems()
}
