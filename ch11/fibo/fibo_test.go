package fibo

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Zero", 0, 0},
		{"First", 1, 1},
		{"Second", 2, 1},
		{"Third", 3, 2},
		{"Fourth", 4, 3},
		{"Fifth", 5, 5},
		{"Tenth", 10, 55},
		{"Negative", -1, -1}, // Проверка обработки ошибки
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fibonacci(tt.input)
			if result != tt.expected {
				t.Errorf("Fibonacci(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFibonacciLarge(t *testing.T) {
	// Тест для большего числа (проверьте на переполнение, если нужно)
	result := Fibonacci(20)
	expected := 6765
	if result != expected {
		t.Errorf("Fibonacci(20) = %d; want %d", result, expected)
	}
}
