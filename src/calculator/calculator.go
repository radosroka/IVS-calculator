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
	GetResult()
}

type result struct {
	value float64
	err   error
}

// Stores current result of equation and slot for selected operation
type SimpleCalc struct {
	result        result
	OperationSlot func(float64, float64) (float64, error)
}

// SimpleCalc constructor
func NewCalc() *SimpleCalc {
	c := new(SimpleCalc)
	c.result = result{
		value: 0,
		err:   nil,
	}
	c.OperationSlot = mathlib.Plus
	return c
}

// Reset calculator to initial state
func (c *SimpleCalc) ClearAll() {
	c.result.value = 0
	c.OperationSlot = mathlib.Plus
}

// Executes selected Operation
func (c *SimpleCalc) Execute(operand float64) {
	c.result.value, c.result.err = c.OperationSlot(c.result.value, operand)
}

// Shows result
func (c *SimpleCalc) Show() {
	fmt.Println(c.result)
}

func (c *SimpleCalc) GetResult() (float64, error) {
	return c.result.value, c.result.err
}
