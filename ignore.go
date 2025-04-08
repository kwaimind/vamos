package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SetupIgnore(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if !strings.HasPrefix(txt, "#") {
			lines = append(lines, txt)
		}
	}
	return lines, scanner.Err()
}
