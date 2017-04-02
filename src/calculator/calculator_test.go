package calculator

import "testing"

var optests = []struct {
	a      float64
	b      float64
	slot   func(*SimpleCalc, float64)
	result float64
}{
	{0, 8, (*SimpleCalc).Plus, 8},
	{0, 8, (*SimpleCalc).Minus, -8},
	{2, 8, (*SimpleCalc).Multiply, 16},
	{4, 2, (*SimpleCalc).Divide, 2},
	{4, 2, (*SimpleCalc).Factorial, 24},
	{2, 2, (*SimpleCalc).Power, 4},
	{16, 4, (*SimpleCalc).NRoot, 2},
	{2, 2, (*SimpleCalc).Mod, 0},
}

func TestOperations(t *testing.T) {
	var calc SimpleCalc

	for _, tt := range optests {
		calc.Result = tt.a
		calc.OperationSlot = tt.slot
		calc.Execute(tt.b)

		if calc.Result != tt.result {
			t.Errorf("Expected %v, but got %v", tt.result, calc.Result)
		}

		calc.ClearAll()
	}
}
