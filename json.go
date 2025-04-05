package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ParseJson(file *os.File) (*PackageJson, error) {
	data := &PackageJson{}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return data, nil
}
