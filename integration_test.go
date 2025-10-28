package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIntegration_FindPackageAndSelectManager(t *testing.T) {
	// Create a temporary workspace
	tmpDir := createTempWorkspace(t)
	defer os.RemoveAll(tmpDir)

	config := InitializeConfig()

	// Test 1: Find package.json in current directory
	packagePath := filepath.Join(tmpDir, "package.json")
	file, err := os.Open(packagePath)
	if err != nil {
		t.Fatalf("Failed to open package.json: %v", err)
	}
	defer file.Close()

	data, err := ParseJson(file)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	packageManager := Select(data)
	if packageManager != config.PNPM {
		t.Errorf("Expected pnpm, got %s", packageManager)
	}

	// Test 2: Find nested package.json by name
	ignoreFiles := []string{"node_modules"}
	foundPath := FindPackageJson(tmpDir, "frontend", ignoreFiles)

	if foundPath == "" {
		t.Fatal("Expected to find frontend package.json")
	}

	if filepath.Base(foundPath) != "package.json" {
		t.Errorf("Expected 'package.json', got '%s'", filepath.Base(foundPath))
	}

	// Verify it's the correct package
	frontendFile, err := os.Open(foundPath)
	if err != nil {
		t.Fatalf("Failed to open frontend package.json: %v", err)
	}
	defer frontendFile.Close()

	frontendData, err := ParseJson(frontendFile)
	if err != nil {
		t.Fatalf("Failed to parse frontend JSON: %v", err)
	}

	if frontendData.Name != "frontend" {
		t.Errorf("Expected name 'frontend', got '%s'", frontendData.Name)
	}

	frontendManager := Select(frontendData)
	if frontendManager != config.Yarn {
		t.Errorf("Expected yarn for frontend, got %s", frontendManager)
	}
}

func TestIntegration_ArgFiltering(t *testing.T) {
	// Test the full arg filtering flow
	args := []string{"-f", "frontend", "test", "--watch"}

	name, filteredArgs := PickFilter(args)

	if name != "frontend" {
		t.Errorf("Expected filter name 'frontend', got '%s'", name)
	}

	expectedArgs := []string{"test", "--watch"}
	if len(filteredArgs) != len(expectedArgs) {
		t.Fatalf("Expected %d args, got %d", len(expectedArgs), len(filteredArgs))
	}

	for i, arg := range expectedArgs {
		if filteredArgs[i] != arg {
			t.Errorf("Expected arg[%d] to be '%s', got '%s'", i, arg, filteredArgs[i])
		}
	}
}

func createTempWorkspace(t *testing.T) string {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "integration-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Create root package.json with pnpm
	rootPkg := `{
		"name": "root",
		"engines": {
			"pnpm": ">=8.9.0"
		}
	}`
	err = os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(rootPkg), 0644)
	if err != nil {
		t.Fatalf("Failed to write root package.json: %v", err)
	}

	// Create frontend subdirectory
	frontendDir := filepath.Join(tmpDir, "packages", "frontend")
	err = os.MkdirAll(frontendDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create frontend dir: %v", err)
	}

	// Create frontend package.json with yarn
	frontendPkg := `{
		"name": "frontend",
		"engines": {
			"yarn": "1.22.0"
		}
	}`
	err = os.WriteFile(filepath.Join(frontendDir, "package.json"), []byte(frontendPkg), 0644)
	if err != nil {
		t.Fatalf("Failed to write frontend package.json: %v", err)
	}

	// Create node_modules to test ignore
	nodeModulesDir := filepath.Join(tmpDir, "node_modules", "some-package")
	err = os.MkdirAll(nodeModulesDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create node_modules: %v", err)
	}

	// Create a package.json inside node_modules that should be ignored
	ignoredPkg := `{
		"name": "ignored-package",
		"engines": {
			"npm": "10.0.0"
		}
	}`
	err = os.WriteFile(filepath.Join(nodeModulesDir, "package.json"), []byte(ignoredPkg), 0644)
	if err != nil {
		t.Fatalf("Failed to write ignored package.json: %v", err)
	}

	return tmpDir
}
