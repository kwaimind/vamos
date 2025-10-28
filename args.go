package main

import (
	"os"
	"strings"
)

func ParseArgs() []string {
	if len(os.Args) < 2 {
		panic("Please provide a command-line argument.")
	}

	return os.Args[1:]
}

func PickFilter(args []string) (string, []string) {

	res := ""
	filteredArgs := []string{}

	skipNext := false
	for i := range len(args) {
		if skipNext {
			skipNext = false
			continue
		}
		if strings.ToLower(args[i]) == "-f" {
			if i+1 < len(args) {
				res = args[i+1]
				skipNext = true
			}
		} else {
			filteredArgs = append(filteredArgs, args[i])
		}
	}

	return res, filteredArgs

}
