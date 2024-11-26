package main

import "math"
import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Set("IsPrime", IsPrime)
}

func IsPrime(value int) int {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return 0
		}
	}
	return value
}
