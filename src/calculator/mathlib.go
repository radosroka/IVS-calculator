package calculator

import (
	"errors"
	"math"
)

// Implementation of mathematical operations

// Plus Return sum of two operands
func Plus(first_op float64, second_op float64) (float64, error) {
	return first_op + second_op, nil
}

// Minus Return difference of two operands
func Minus(first_op float64, second_op float64) (float64, error) {
	return first_op - second_op, nil
}

// Multiply Return product of two operands
func Multiply(first_op float64, second_op float64) (float64, error) {
	return first_op * second_op, nil
}

// Divide Return quotient of two operands
func Divide(first_op float64, second_op float64) (float64, error) {
	if second_op == 0 {
		return 0, errors.New("Division by zero")
	}
	return first_op / second_op, nil
}

// Factorial Return factorial of first operand, the second argument just for
// compatibility with OperationSlot interface
func Factorial(first_op float64, second_op float64) (float64, error) {
	var res, k int64
	res = 1
	if float64(int64(first_op)) != first_op || first_op < 0 {
		return 0, errors.New("Factorial of non integer or negative value can't be calculated")
	}
	if first_op > 25 {
		return 0, errors.New("Integer overflow")
	}
	input := int64(first_op)
	for k = 1; k <= input; k++ {
		res *= k
	}
	return float64(res), nil
}

// Power Return first operand (base) taken to the power of
// the second_operand (exponent)
func Power(first_op float64, second_op float64) (float64, error) {
	return math.Pow(first_op, second_op), nil
}

// NRoot Return a number used n times (second operand) in a multiplication
// gives first operand
func NRoot(first_op float64, second_op float64) (float64, error) {
	return math.Pow(first_op, 1.0/second_op), nil
}

// Mod Return remainder of first operand divided by the second operand
func Mod(first_op float64, second_op float64) (float64, error) {
	return math.Mod(first_op, second_op), nil
}
