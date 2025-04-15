package main

import (
	"encoding/json"
	"io"
	"os"
)

func ParseJson(file *os.File) (*PackageJson, error) {
	data := &PackageJson{}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
