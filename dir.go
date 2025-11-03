package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func FindPackageJson(dir string, packageName string, ignoreFiles []string) (string, error) {
	var result string

	err := filepath.WalkDir(dir,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			isRoot := path == dir
			shouldIgnore := slices.Contains(ignoreFiles, d.Name())
			skipDotFiles := strings.HasPrefix(d.Name(), ".")

			if !isRoot && d.IsDir() && (shouldIgnore || skipDotFiles) {
				return filepath.SkipDir
			}

			if d.Name() == packageJSONName {

				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				data, err := ParseJson(file)
				if err != nil {
					return err
				}

				if data.Name == packageName {
					result = path
					return filepath.SkipAll
				}
			}

			return nil
		})

	return result, err
}
