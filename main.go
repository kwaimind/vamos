package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	args := ParseArgs()
	config := InitializeConfig()

	name := PickFilter(args)

	packagePath := config.PackagejsonName

	if name != "" {
		packagePath = FindPackageJson(config.RootDir, name)
		fmt.Println(name, packagePath)
	}

	file, err := os.Open(packagePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := ParseJson(file)
	if err != nil {
		panic(err)
	}

	packageManager := Select(data)

	nextargs := args

	if packageManager == config.NPM {
		nextargs = append([]string{config.NPMRun}, args...)
	}

	cmd := exec.Command(packageManager, nextargs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
