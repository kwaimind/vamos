package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FindPackageJson(rootDir string, packageName string) string {

	config := InitializeConfig()

	var result string

	err := filepath.Walk(rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Name() == config.PackagejsonName {

				file, err := os.Open(path)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()

				data, err := ParseJson(file)
				if err != nil {
					return err
				}

				if data.Name == packageName {
					result = path
					return fmt.Errorf("found")
				}
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return result
}
