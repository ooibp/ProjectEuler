package main

import "fmt"

func fibo(first, second int64) int64 {
	return first + second
}

func main() {
	x, y := int64(1), int64(1)
	var dummy, sum int64
	for y <= 4000000 {
		dummy = fibo(int64(x), int64(y))
		if dummy%2 == 0 {
			sum += dummy
		}
		x = y
		y = dummy
	}
	fmt.Println(sum)
}
