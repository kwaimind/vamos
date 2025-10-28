package main

import (
	"testing"
)

func TestCapitalizeFirst_NormalString(t *testing.T) {
	result := CapitalizeFirst("hello world")
	expected := "Hello world"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestCapitalizeFirst_AlreadyCapitalized(t *testing.T) {
	result := CapitalizeFirst("Hello world")
	expected := "Hello world"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestCapitalizeFirst_EmptyString(t *testing.T) {
	result := CapitalizeFirst("")
	expected := ""

	if result != expected {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestCapitalizeFirst_SingleChar(t *testing.T) {
	result := CapitalizeFirst("a")
	expected := "A"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestCapitalizeFirst_WithUnicode(t *testing.T) {
	result := CapitalizeFirst("éclair")
	expected := "Éclair"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestCapitalizeFirst_AllLowercase(t *testing.T) {
	result := CapitalizeFirst("test string")
	expected := "Test string"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
