package main

import "fmt"

func main() {
	var longest int64
	for index := 1000000; index > 1; index-- {
		num := index
		var dummy int64

		for num != 1 {
			dummy++
			if num%2 == 0 {
				num /= 2
			} else {
				num = 3*num + 1
			}
		}

		if dummy > longest {
			longest = dummy
			fmt.Println(index)
			fmt.Println(longest)
		}
	}
}

// The following iterative sequence is defined for the set of positive integers:
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
// Using the rule above and starting with 13, we generate the following sequence:
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
// Which starting number, under one million, produces the longest chain?
// NOTE: Once the chain starts the terms are allowed to go above one million.
