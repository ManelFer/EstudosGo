// aqui vai ser o codigo de teste para a calculadora simples

package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(5, 3)
	if result != 8 {
		t.Errorf("Expected 8, got %v", result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, got %v", result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(5, 3)
	if result != 15 {
		t.Errorf("Expected 15, got %v", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := divide(5, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5/3 {
		t.Errorf("Expected 5/3, got %v", result)
	}
}
