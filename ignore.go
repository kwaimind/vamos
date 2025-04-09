package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func SetupIgnore(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		isDir, _ := isDirectory(txt)
		if !strings.HasPrefix(txt, "#") && isDir {
			lines = append(lines, txt)
		}
	}

	return lines, scanner.Err()
}
