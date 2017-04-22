package main

import (
	"fmt"
	"os"
	"strconv"
	"mathlib"
)

const MAX = 1000

func main() {
	var array[MAX] float64

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Bad arguments\n")
		return
	}

	N, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < N && i < MAX; i++ {
		fmt.Scanf("%f", &array[i])
//		fmt.Printf("%d\n", array[i])
	}

	var res float64 = 0.0
	var mean float64 = 0.0

	for i := 0; i < N && i < MAX; i++ {
		h, _ := mathlib.Plus(float64(i), 1.0)
		f, _ := mathlib.Divide(1.0, h)
  		d, _ := mathlib.Minus(array[i], mean)
  		dd, _ := mathlib.Multiply(d, f)
  		meann, _ := mathlib.Plus(mean, dd)
  		mean = meann
  		c, _ := mathlib.Minus(1.0, f)
  		x, _ := mathlib.Multiply(dd, d)
  		y, _ := mathlib.Plus(res, x)
  		z, _ := mathlib.Multiply(c, y)
  		ress, _ := mathlib.NRoot(z, 2)
  		res = ress
	}

	fmt.Printf("Deviation is -- %v", res)

}