package main

import "fmt"

func main() {
	total := 0
	for index := 0; index <= 100; index++ {
		if index%3 == 0 || index%5 == 0 {
			total += index
		}
	}
	fmt.Println("Sum = ", total)
}
