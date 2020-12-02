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
	return a
}

func getDivisorsVal(h int, eCount []int) int {
	var div []int
	x := int(math.Sqrt(float64(h)))
	for i := 1; i <= x; i++ {
		if h%i == 0 {
			j := h / i

			if j != i {
				if eCount[j] < 50 {
					eCount[j] += 1
					div = append(div, j)
				}
			}

			if eCount[i] < 50 {
				eCount[i] += 1
				div = append(div, i)
			}
		}
	}

	val := addArray(div) * 11
	return val
}

func main() {
	var elfCount []int
	var housesDelivered []int
	target := 29000000
	h := 0
	elfCount = append(elfCount, 0)
	housesDelivered = append(housesDelivered, 0)

	for true {
		h += 1
		elfCount = append(elfCount, 0)
		val := getDivisorsVal(h, elfCount)
		housesDelivered = append(housesDelivered, val)

		if val >= target {
			fmt.Printf("\nFound it!\n\tHouse: %d Ans: %d\n", h, val)
			os.Exit(0)
		}
	}
}
