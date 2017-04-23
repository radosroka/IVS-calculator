package main

import (
	"calculator"
	"fmt"
	"github.com/pkg/profile"
	"os"
	"runtime"
	"strconv"
	"time"
)

const MAX = 1000

func main() {
	time.Sleep(time.Second*2)
	fmt.Println("Starting profile...")
	runtime.SetCPUProfileRate(10000000000)
	p := profile.Start(profile.CPUProfile, profile.ProfilePath("./profiling/"))
	var array[MAX] float64

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Bad arguments\n")
		return
	}

	N, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < N && i < MAX; i++ {
		fmt.Scanf("%f", &array[i])
	}
	var res float64 = 0.0
	var mean float64 = 0.0

	for i := 0; i < N && i < MAX; i++ {
		h, _ := calculator.Plus(float64(i), 1.0)
		f, _ := calculator.Divide(1.0, h)
		d, _ := calculator.Minus(array[i], mean)
		dd, _ := calculator.Multiply(d, f)
		meann, _ := calculator.Plus(mean, dd)
		mean = meann
		c, _ := calculator.Minus(1.0, f)
		x, _ := calculator.Multiply(dd, d)
		y, _ := calculator.Plus(res, x)
		z, _ := calculator.Multiply(c, y)
		ress, _ := calculator.NRoot(z, 2)
		res = ress
	}

	fmt.Printf("Deviation is -- %v\n", res)
	p.Stop()
}
