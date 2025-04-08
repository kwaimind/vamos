package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	args := ParseArgs()
	config := InitializeConfig()

	name := PickFilter(args)

	packagePath := config.PackagejsonName
	ignoreFile := config.GitIgnore

	if name != "" {
		nextPackagePath := FindPackageJson(config.RootDir, name)
		packagePath = nextPackagePath
		ignoreFile = strings.Replace(nextPackagePath, config.PackagejsonName, config.GitIgnore, 1)
	}

	ignoreFiles, _ := SetupIgnore(ignoreFile)

	fmt.Println(ignoreFiles)

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

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdErr := cmd.Run()
	if cmdErr != nil {
		panic(cmdErr)
	}
}
