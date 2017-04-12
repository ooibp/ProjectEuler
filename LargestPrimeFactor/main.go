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
	var num int64 = 600851475143
	for index := int64(sqrt(num)); index >= 1; index-- {
		if num%index == 0 && isprime(index) {
			fmt.Println(index)
			break
		}
	}
}

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
