package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
