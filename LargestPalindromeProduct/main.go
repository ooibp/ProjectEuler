package main

import "fmt"

func main() {
	var a, b int64 = 999, 999
	var dummy, first, second, third, fourth, fifth, sixth int64
	var largest []int64

	for index := a; index >= 100; index-- {
		for jndex := b; jndex >= 100; jndex-- {
			dummy = index * jndex
			first = dummy / 100000
			dummy = dummy % 100000
			second = dummy / 10000
			dummy = dummy % 10000
			third = dummy / 1000
			dummy = dummy % 1000
			fourth = dummy / 100
			dummy = dummy % 100
			fifth = dummy / 10
			sixth = dummy % 10

			if first == sixth && second == fifth && third == fourth {
				largest = append(largest, index*jndex)
			}
		}
	}

	var large int64
	for _, value := range largest {
		if value > large {
			large = value
		}
	}
	fmt.Println(large)
}

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
