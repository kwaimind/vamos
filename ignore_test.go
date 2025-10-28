package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSetupIgnore_ValidFile(t *testing.T) {
	content := `node_modules
dist/
.env
# This is a comment
build

.git/`

	tmpFile := createTempIgnoreFile(t, content)
	defer os.Remove(tmpFile)

	result, err := SetupIgnore(tmpFile)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{"node_modules", "dist", ".env", "build", ".git"}
	if len(result) != len(expected) {
		t.Fatalf("Expected %d entries, got %d", len(expected), len(result))
	}

	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected result[%d] to be '%s', got '%s'", i, val, result[i])
		}
	}
}

func TestSetupIgnore_EmptyFile(t *testing.T) {
	content := ``

	tmpFile := createTempIgnoreFile(t, content)
	defer os.Remove(tmpFile)

	result, err := SetupIgnore(tmpFile)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}

func TestSetupIgnore_OnlyComments(t *testing.T) {
	content := `# Comment 1
# Comment 2
# Comment 3`

	tmpFile := createTempIgnoreFile(t, content)
	defer os.Remove(tmpFile)

	result, err := SetupIgnore(tmpFile)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}

func TestSetupIgnore_WithWhitespace(t *testing.T) {
	content := `  node_modules

dist
   build   `

	tmpFile := createTempIgnoreFile(t, content)
	defer os.Remove(tmpFile)

	result, err := SetupIgnore(tmpFile)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{"node_modules", "dist", "build"}
	if len(result) != len(expected) {
		t.Fatalf("Expected %d entries, got %d", len(expected), len(result))
	}

	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected result[%d] to be '%s', got '%s'", i, val, result[i])
		}
	}
}

func TestSetupIgnore_FileNotFound(t *testing.T) {
	result, err := SetupIgnore("/nonexistent/path/.gitignore")

	if err == nil {
		t.Fatal("Expected error for non-existent file, got nil")
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result on error, got %v", result)
	}
}

func createTempIgnoreFile(t *testing.T, content string) string {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "ignoretest-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	tmpFile := filepath.Join(tmpDir, ".gitignore")
	err = os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	return tmpFile
}
