package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPackageJson_FindsMatch(t *testing.T) {
	tmpDir := createTempPackage(t, "my-project")
	defer os.RemoveAll(tmpDir)

	result := FindPackageJson(tmpDir, "my-project", make([]string, 0))

	if result == "" {
		t.Fatal("Expected to find package.json, got empty string")
	}
	if filepath.Base(result) != "package.json" {
		t.Errorf("Expected file named 'package.json', got '%s'", filepath.Base(result))
	}
}

func TestFindPackageJson_NoMatch(t *testing.T) {
	tmpDir := createTempPackage(t, "not-the-right-name")
	defer os.RemoveAll(tmpDir)

	result := FindPackageJson(tmpDir, "target-name", make([]string, 0))

	if result != "" {
		t.Errorf("Expected no match, got '%s'", result)
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
