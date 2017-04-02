package main

import "./calculator"

func main() {
	calc := calculator.SimpleCalc{Result: 0.0, OperationSlot: (*calculator.SimpleCalc).Plus}
	calc.Execute(8.0)
	calc.OperationSlot = (*calculator.SimpleCalc).Minus
	calc.Execute(3.0)
	calc.Show()
	calc.OperationSlot = (*calculator.SimpleCalc).Factorial
	calc.Execute(0.0)
	calc.Show()
	calc.OperationSlot = (*calculator.SimpleCalc).Power
	calc.Execute(2)
	calc.Show()
	calc.ClearAll()
	calc.Execute(29)
	calc.OperationSlot = (*calculator.SimpleCalc).Mod
	calc.Execute(5)
	calc.Show()
	calc.OperationSlot = (*calculator.SimpleCalc).NRoot
	calc.Execute(2)
	calc.Show()
}
