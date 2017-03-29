package calculator

import "fmt"
import "math"

// Calculator defines Operations of calculator
type Calculator interface {
    Plus(Calculator, float64)
    Minus(Calculator, float64)
    Multiply(Calculator, float64)
    Divide(Calculator, float64)
    Factorial(Calculator, float64)
    Power(Calculator, float64)
    NRoot(Calculator, float64)
    Mod(Calculator, float64)
    ClearAll(Calculator)
    Execute(Calculator, float64)
    Show(Calculator)
}

// Stores current Result of equation and slot for selected Operation
type SimpleCalc struct {
    Result float64
    OperationSlot func(*SimpleCalc, float64)
}

// Implementation of mathematical Operations
func(calc *SimpleCalc) Plus(operand float64) {
    calc.Result += operand
}

func(calc *SimpleCalc) Minus(operand float64) {
    calc.Result -= operand
}

func(calc *SimpleCalc) Multiply(operand float64) {
    calc.Result *= operand
}

// TODO Division by zero??
func(calc *SimpleCalc) Divide(operand float64) {
    calc.Result /= operand
}

// Argument just for OperationSlot compatibility 
// TODO int overflow
// TODO math error if calc.Result can't be casted to Int
func(calc *SimpleCalc) Factorial(operand float64) {
    res := 1
    input := int(calc.Result)
    for k := 1; k <= input; k++ {
        res *= k
    }
    calc.Result = float64(res)
}

func(calc *SimpleCalc) Power(operand float64) {
    calc.Result = math.Pow(calc.Result, operand)
}

func(calc *SimpleCalc) NRoot(operand float64) {
    calc.Result = math.Pow(calc.Result, 1.0 / operand)
}

func(calc *SimpleCalc) Mod(operand float64) {
    calc.Result = math.Mod(calc.Result, operand)
}

// Reset calculator to initial state
func(calc *SimpleCalc) ClearAll() {
    calc.Result = 0;
    calc.OperationSlot = (*SimpleCalc).Plus
}

// Executes selected Operation
func(calc *SimpleCalc) Execute(operand float64) {
    calc.OperationSlot(calc, operand)
}

// Shows Result
func(calc *SimpleCalc) Show() {
    fmt.Println(calc.Result)
}
