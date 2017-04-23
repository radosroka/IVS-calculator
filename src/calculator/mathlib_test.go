package calculator

import "testing"

var optests = []struct {
	a      float64
	b      float64
	op     func(float64, float64) (float64, error)
	result float64
}{
	{0, 8, Plus, 8},
	{0, 8, Minus, -8},
	{2, 8, Multiply, 16},
	{4, 2, Divide, 2},
	{4, 2, Factorial, 24},
	{2, 2, Power, 4},
	{16, 4, NRoot, 2},
	{2, 2, Mod, 0},
}

func TestOperations(t *testing.T) {

	for _, tt := range optests {
		if res, err := tt.op(tt.a, tt.b); res != tt.result {
			t.Errorf("Expected %v, but got %v", tt.result, res)
		} else if err != nil {
			t.Error(err)
		}
	}
}

func TestDivisionByZero(t *testing.T) {
	if _, err := Divide(10, 0); err == nil {
		t.Error("Division by zero should be not possible")
	}
}

func TestFactorial(t *testing.T) {
	if _, err := Factorial(1.5, 8); err.Error() != "Factorial of non integer or negative value can't be calculated" {
		t.Error("Expected error when trying to factor non-interger or negative value")
	}

	if _, err := Factorial(-2, 8); err.Error() != "Factorial of non integer or negative value can't be calculated" {
		t.Error("Expected error when trying to factor non-interger or negative value")
	}

	if _, err := Factorial(26, 8); err.Error() != "Integer overflow" {
		t.Error("Expected integer overflow")
	}
}
