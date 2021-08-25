package main

import (
	"fmt"
	"log"

	"github.com/disiqueira/gotree"
	"github.com/urfave/cli/v2"
)

func PrintTree(_ *cli.Context) error {
	if len(CurrentData) == 0 {
		log.Println("This file does not contain any data.")
		return nil
	}

	root := gotree.New(InputFilePath)
	processNode(&root, CurrentData)
	log.Println(root.Print())
	return nil
}

func processNode(parent *gotree.Tree, data map[string]interface{}) {
	for key, value := range data {
		if mapValue, ok := value.(map[string]interface{}); ok {
			child := (*parent).Add(key)
			processNode(&child, mapValue)
		} else {
			(*parent).Add(fmt.Sprintf("%s: %s", key, value))
		}
	}
}
