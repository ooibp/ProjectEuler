package main

import (
	"fmt"
	"math"
)

func main() {
	var num []int64

	sqrt := func(a int64) int64 {
		return int64(math.Sqrt(float64(a)))
	}

	isprime := func(b int64) bool {
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

	for index := int64(2); len(num) < 10001; index++ {
		if isprime(index) {
			num = append(num, index)
		}
	}
	fmt.Println(num[10000])
}

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
