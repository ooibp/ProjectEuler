package main

import "fmt"

func main() {
	sum := func(a []int64) int64 {
		total := int64(0)
		for _, value := range a {
			total += value
		}
		return total
	}

	num := make([]int64, 100)
	square := make([]int64, 100)

	for index := int64(1); index <= 100; index++ {
		num[index-1] = index
		square[index-1] = index * index
	}

	total := sum(num)
	difference := total*total - sum(square)
	fmt.Println(difference)
}

// The sum of the squares of the first ten natural numbers is
// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
