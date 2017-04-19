package mathlib

import (
	"errors"
	"math"
)

// Implementation of mathematical operations
func Plus(first_op float64, second_op float64) (float64, error) {
	return first_op + second_op, nil
}

func Minus(first_op float64, second_op float64) (float64, error) {
	return first_op - second_op, nil
}

func Multiply(first_op float64, second_op float64) (float64, error) {
	return first_op * second_op, nil
}

func Divide(first_op float64, second_op float64) (float64, error) {
	if second_op == 0 {
		return 0, errors.New("Division by zero")
	}
	return first_op / second_op, nil
}

// Argument just for OperationSlot compatibility
// TODO int overflow
func Factorial(first_op float64, second_op float64) (float64, error) {
	res := 1
	if float64(int(first_op)) != first_op {
		return 0, errors.New("Factorial of non integer value can't be calculated")
	}
	input := int(first_op)
	for k := 1; k <= input; k++ {
		res *= k
	}
	return float64(res), nil
}

func Power(first_op float64, second_op float64) (float64, error) {
	return math.Pow(first_op, second_op), nil
}

func NRoot(first_op float64, second_op float64) (float64, error) {
	return math.Pow(first_op, 1.0/second_op), nil
}

func Mod(first_op float64, second_op float64) (float64, error) {
	return math.Mod(first_op, second_op), nil
}
