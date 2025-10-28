package main

import (
	"bufio"
	"os"
	"strings"
)

func SetupIgnore(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if txt != "" && !strings.HasPrefix(txt, "#") {
			// Remove trailing slashes from directory entries
			txt = strings.TrimSuffix(txt, "/")
			lines = append(lines, txt)
		}
	}

	return lines, scanner.Err()
}
