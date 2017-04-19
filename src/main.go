package main

import (
	"./calculator"
	"./mathlib"
)

func main() {
	calc := calculator.SimpleCalc{
		Result: calculator.Result{
			Value:    0,
			Err:      nil,
			Negative: false,
		},
		OperationSlot: mathlib.Plus,
	}
	calc.Execute(8.0)
	calc.OperationSlot = mathlib.Minus
	calc.Execute(3.0)
	calc.Show()
	calc.OperationSlot = mathlib.Factorial
	calc.Execute(0.0)
	calc.Show()
	calc.OperationSlot = mathlib.Power
	calc.Execute(2)
	calc.Show()
	calc.ClearAll()
	calc.Execute(29)
	calc.OperationSlot = mathlib.Mod
	calc.Execute(5)
	calc.Show()
	calc.OperationSlot = mathlib.NRoot
	calc.Execute(2)
	calc.Show()
	calc.OperationSlot = mathlib.Divide
	calc.Execute(0)
	calc.Show()
}
