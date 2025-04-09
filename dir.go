package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func FindPackageJson(rootDir string, packageName string, ignoreFiles []string) string {

	config := InitializeConfig()

	var result string

	err := filepath.Walk(rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			isRoot := path == rootDir
			shouldIgnore := slices.Contains(ignoreFiles, info.Name())
			skipDotFiles := strings.HasPrefix(info.Name(), ".")

			if !isRoot && info.IsDir() && (shouldIgnore || skipDotFiles) {
				return filepath.SkipDir
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
					return nil
				}
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return result
}
