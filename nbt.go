package main

import (
	"bytes"
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

	defer file.Close()

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
		if OutputFilePath == "" {
			OutputFilePath = InputFilePath
		}

		err := LoadData()
		if err != nil {
			return err
		}

		return actionFunc(context)
	}
}

func SaveData() error {
	data, err := nbt.Marshal(CurrentData)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(OutputFilePath, buffer.Bytes(), 0644)
}
