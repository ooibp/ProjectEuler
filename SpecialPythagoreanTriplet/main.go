package main

import "fmt"

func main() {
	for index := 2; index < 1000; index++ {
		for jndex := 2; jndex < 1000; jndex++ {
			for kndex := 2; kndex < 1000; kndex++ {
				if index*index == jndex*jndex+kndex*kndex {
					if index+jndex+kndex == 1000 {
						fmt.Println(index, jndex, kndex)
						fmt.Println(index * jndex * kndex)
					}
				}
			}
		}
	}
}

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
