package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/yarlson/pin"
)

func main() {

	p := pin.New("⚡️ Vamos...",
		pin.WithSpinnerColor(pin.ColorCyan),
		pin.WithTextColor(pin.ColorYellow),
	)
	cancel := p.Start(context.Background())
	defer cancel()

	args := ParseArgs()
	config := InitializeConfig()

	name := PickFilter(args)

	packagePath := config.PackagejsonName

	if name != "" {
		ignoreFile := config.GitIgnore
		ignoreFiles, _ := SetupIgnore(ignoreFile)
		nextPackagePath := FindPackageJson(config.RootDir, name, ignoreFiles)
		packagePath = nextPackagePath
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

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdErr := cmd.Run()
	if cmdErr != nil {
		panic(cmdErr)
	}
}
