package main

import (
	"fmt"
	"math/big"
)

func main() {
	var fact40, fact20 big.Int
	fact40.MulRange(1, 40)
	fact20.MulRange(1, 20)
	fact40.Div(&fact40, &fact20)
	fact40.Div(&fact40, &fact20)
	fmt.Println(fact40)
}

// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
// How many such routes are there through a 20×20 grid?
