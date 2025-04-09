package main

import (
	"os"
	"testing"
)

func TestParseJson_Valid(t *testing.T) {
	jsonContent := `{
		"name": "my-project",
		"engines": {
			"npm": "10.9.2"
		},
		"scripts": {
			"test": "echo 'running npm'"
		}
	}`
	file := tempFile(t, jsonContent)
	defer os.Remove(file.Name())

	pkg, err := ParseJson(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pkg.Name != "my-project" {
		t.Errorf("Expected name to be 'my-project', got %s", pkg.Name)
	}
	if pkg.Engines.NPM != "10.9.2" {
		t.Errorf("Expected node version to be '10.9.2', got %s", pkg.Engines.NPM)
	}

}

func TestParseJson_InvalidJSON(t *testing.T) {
	jsonContent := `{"name": "my-package", "version":` // Invalid JSON
	file := tempFile(t, jsonContent)
	defer os.Remove(file.Name())

	_, err := ParseJson(file)
	if err == nil {
		t.Fatal("Expected error for invalid JSON, got nil")
	}
}

func tempFile(t *testing.T, content string) *os.File {
	t.Helper()
	tmp, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err := tmp.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if _, err := tmp.Seek(0, 0); err != nil {
		t.Fatalf("Failed to seek in temp file: %v", err)
	}
	return tmp
}
