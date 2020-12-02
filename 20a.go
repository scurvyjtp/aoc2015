package main

import (
	"fmt"
	"math"
	"os"
)

func getProduct(x int, h int) int {
	return (h / x) * (x * 10)
}

func addArray(div []int) int {
	var a int
	for _, d := range div {
		a += d
	}
	//fmt.Printf("  %s : %d\n", div, a)
	return a
}

func getDivisorsVal(h int) int {
	var div []int
	x := int(math.Sqrt(float64(h)))
	for i := 1; i <= x; i++ {
		if h%i == 0 {
			if (h / i) == i {
				div = append(div, i)
			} else {
				div = append(div, i, (h / i))
			}
		}
	}

	val := addArray(div) * 10
	return val
}

func main() {
	var houses []int
	target := 29000000

	house := 0
	houses = append(houses, 0)

	for true {
		house += 1
		val := getDivisorsVal(house)
		houses = append(houses, val)

		if val >= target {
			fmt.Printf("\nFound it!\n\tHouse: %d Ans: %d\n", house, val)
			os.Exit(0)
		}
	}
}
