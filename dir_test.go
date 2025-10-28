package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPackageJson_FindsMatch(t *testing.T) {
	config := InitializeConfig()
	tmpDir := createTempPackage(t, "my-project")
	defer os.RemoveAll(tmpDir)

	result, err := FindPackageJson(tmpDir, "my-project", make([]string, 0), config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result == "" {
		t.Fatal("Expected to find package.json, got empty string")
	}
	if filepath.Base(result) != "package.json" {
		t.Errorf("Expected file named 'package.json', got '%s'", filepath.Base(result))
	}
}

func TestFindPackageJson_NoMatch(t *testing.T) {
	config := InitializeConfig()
	tmpDir := createTempPackage(t, "not-the-right-name")
	defer os.RemoveAll(tmpDir)

	result, err := FindPackageJson(tmpDir, "target-name", make([]string, 0), config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "" {
		t.Errorf("Expected no match, got '%s'", result)
	}
}

func TestFindPackageJson_WithIgnore(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "pkgtest-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a package in node_modules that should be ignored
	nodeModulesDir := filepath.Join(tmpDir, "node_modules")
	err = os.MkdirAll(nodeModulesDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create node_modules: %v", err)
	}

	jsonContent := `{"name": "should-be-ignored"}`
	err = os.WriteFile(filepath.Join(nodeModulesDir, "package.json"), []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Search with node_modules in ignore list
	config := InitializeConfig()
	ignoreList := []string{"node_modules"}
	result, err := FindPackageJson(tmpDir, "should-be-ignored", ignoreList, config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should not find it because it's ignored
	if result != "" {
		t.Errorf("Expected no match (ignored), got '%s'", result)
	}
}

func TestFindPackageJson_SkipDotFiles(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "pkgtest-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a package in .hidden directory that should be skipped
	hiddenDir := filepath.Join(tmpDir, ".hidden")
	err = os.MkdirAll(hiddenDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create .hidden: %v", err)
	}

	jsonContent := `{"name": "hidden-package"}`
	err = os.WriteFile(filepath.Join(hiddenDir, "package.json"), []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	config := InitializeConfig()
	result, err := FindPackageJson(tmpDir, "hidden-package", make([]string, 0), config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should not find it because dot directories are skipped
	if result != "" {
		t.Errorf("Expected no match (dot directory), got '%s'", result)
	}
}

func TestFindPackageJson_WithInvalidJSON(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "pkgtest-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a package.json with invalid JSON
	invalidJSON := `{"name": "test", invalid json}`
	err = os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(invalidJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	config := InitializeConfig()
	result, err := FindPackageJson(tmpDir, "test", make([]string, 0), config)

	// Should return error because JSON parsing failed
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
	if result != "" {
		t.Errorf("Expected empty result on error, got '%s'", result)
	}
}

func createTempPackage(t *testing.T, name string) string {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "pkgtest-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	jsonContent := `{
		"name": "` + name + `",
		"engines": {
			"npm": "10.9.2"
		},
		"scripts": {
			"test": "echo test"
		}
	}`

	err = os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	return tmpDir
}
