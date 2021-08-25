package main

import (
	"compress/gzip"
	"io/ioutil"
	"os"

	"github.com/Tnze/go-mc/nbt"
	"github.com/urfave/cli/v2"
)

var CurrentData map[string]interface{}

func LoadData() error {
	file, err := os.Open(InputFilePath)
	if err != nil {
		return err
	}

	reader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return nbt.Unmarshal(data, &CurrentData)
}

func FileAction(actionFunc cli.ActionFunc) cli.ActionFunc {
	return func(context *cli.Context) error {
		err := LoadData()
		if err != nil {
			return err
		}

		return actionFunc(context)
	}
}
