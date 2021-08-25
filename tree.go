package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/disiqueira/gotree"
	"github.com/urfave/cli/v2"
)

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

func PrintTree(_ *cli.Context) error {
	if len(CurrentData) == 0 {
		log.Println("This file does not contain any data.")
		return nil
	}

	root := gotree.New(InputFilePath)
	processNode(&root, CurrentData)
	log.Println(printTree(root))
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

func printTree(tree gotree.Tree) string {
	return tree.Text() + newLine + printItems(tree.Items(), []bool{})
}

func printItems(tree []gotree.Tree, spaces []bool) string {
	var result string
	for i, f := range tree {
		last := i == len(tree)-1
		result += printText(f.Text(), spaces, last)
		if len(f.Items()) > 0 {
			spacesChild := append(spaces, last)
			result += printItems(f.Items(), spacesChild)
		}
	}
	return result
}

func printText(text string, spaces []bool, last bool) string {
	var result string
	for _, space := range spaces {
		if space {
			result += emptySpace
		} else {
			result += continueItem
		}
	}

	indicator := middleItem
	if last {
		indicator = lastItem
	}

	var out string
	lines := strings.Split(text, "\n")
	for i := range lines {
		text := lines[i]
		if i == 0 {
			out += result + indicator + text + newLine
			continue
		}
		if last {
			indicator = emptySpace
		} else {
			indicator = continueItem
		}
		out += result + indicator + text + newLine
	}

	return out
}
