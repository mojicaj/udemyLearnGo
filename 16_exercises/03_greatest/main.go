package main

import "fmt"

func main() {
	n := greatest(15, 4, 10, 8)
	fmt.Println(n)
}

func greatest(si ...int) int {
	var big int
	for _, v := range si {
		if v > big {
			big = v
		}
	}
	return big
}
