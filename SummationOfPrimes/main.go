package main

import (
	"fmt"
	"math"
)

func sqrt(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

func isprime(b int64) bool {
	if b <= 1 {
		return false
	}
	for index := int64(sqrt(b)); index >= 1; index-- {
		if index == 1 {
			return true
		}
		if b%index == 0 {
			return false
		}
	}
	return true
}

func main() {
	var total int64
	for index := int64(2); index <= 2000000; index++ {
		if isprime(index) {
			total += index
		}
	}
	fmt.Println(total)
}
