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

func PickFilter(args []string) string {

	res := ""

	for i := range len(args) - 1 {
		if strings.ToLower(args[i]) == "-f" {
			res = args[i+1]

		}
	}

	return res

}
