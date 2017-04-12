package main

import "fmt"

func sum(a []int64) int64 {
	total := int64(0)
	for _, value := range a {
		total += value
	}
	return total
}

func main() {
	divisor := []int64{11, 13, 14, 15, 16, 17, 18, 19}
	for index := int64(1); ; index++ {
		var remain []int64
		for _, jndex := range divisor {
			remain = append(remain, index%jndex)
		}

		if sum(remain) == 0 {
			fmt.Println(index)
			break
		}
	}
}

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
