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
