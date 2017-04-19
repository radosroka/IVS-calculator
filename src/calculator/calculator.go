package calculator

import (
	"fmt"
	"mathlib"
)

// Calculator defines Operations of calculator
type Calculator interface {
	ClearAll()
	Execute(float64)
	Show()
}

type Result struct {
	Value    float64
	Err      error
	Negative bool
}

// Stores current Result of equation and slot for selected Operation
type SimpleCalc struct {
	Result        Result
	OperationSlot func(float64, float64) (float64, error)
}

// Reset calculator to initial state
func (ca *SimpleCalc) ClearAll() {
	ca.Result.Value = 0
	ca.Result.Negative = false
	ca.OperationSlot = mathlib.Plus
}

// Executes selected Operation
func (ca *SimpleCalc) Execute(operand float64) {
	ca.Result.Value, ca.Result.Err = ca.OperationSlot(ca.Result.Value, operand)
}

// Shows Result
func (calc *SimpleCalc) Show() {
	fmt.Println(calc.Result)
}
