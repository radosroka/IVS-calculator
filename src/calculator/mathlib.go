package calculator

import (
	"errors"
	"math"
)

// Implementation of mathematical operations

// Plus Return sum of two operands
func Plus(a, b float64) (float64, error) {
	return a + b, nil
}

// Minus Return difference of two operands
func Minus(a, b float64) (float64, error) {
	return a - b, nil
}

// Multiply Return product of two operands
func Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

// Divide Return quotient of two operands
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

// Factorial Return factorial of first operand, the second argument just for
// compatibility with OperationSlot interface
func Factorial(a, b float64) (float64, error) {
	var res, k int64
	res = 1
	if float64(int64(a)) != a || a < 0 {
		return 0, errors.New("Factorial of non integer or negative value can't be calculated")
	}
	if a > 25 {
		return 0, errors.New("Integer overflow")
	}
	input := int64(a)
	for k = 1; k <= input; k++ {
		res *= k
	}
	return float64(res), nil
}

// Power Return first operand (base) taken to the power of
// the berand (exponent)
func Power(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}

// NRoot Return a number used n times (second operand) in a multiplication
// gives first operand
func NRoot(a, b float64) (float64, error) {
	return math.Pow(a, 1.0/b), nil
}

// Mod Return remainder of first operand divided by the second operand
func Mod(a, b float64) (float64, error) {
	return math.Mod(a, b), nil
}
