package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	dummy := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()

	this := func(s string) int {
		total := 0
		for _, value := range s {
			num, err := strconv.Atoi(string(value))
			if err != nil {
				panic(err)
			}
			total += num
		}
		return total
	}(dummy)
	fmt.Println(this)
}

// 215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
// What is the sum of the digits of the number 2^1000?
