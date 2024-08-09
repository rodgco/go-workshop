package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello young gopher")

	a := "double"
	b := "rainbow"
	// TODO print `a` and `b` to the console as part of a sentence
	fmt.Printf("a is %v and b %v.\n", a, b)

	const c = 42
	const d = 6.022e23
	// TODO use math.Min to find the minimum of the two constants `c` and `d` and print the result to the console
	fmt.Printf("Minimum is %v.\n", math.Min(c, d))

	// TODO add a variable `e` to make the following console message return true
	const powerLevel = 9000
	e := 10_000
	fmt.Printf("It's over 9000? %v\n", e > powerLevel)
}
