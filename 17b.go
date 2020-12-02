package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []int {
	var l []int
	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		l = append(l, n)
	}

	file.Close()
	return l
}

func arrAdd(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func subSum(nums []int, target int, allC *[][]int, part ...[]int) int {
	total := 0
	x := 0
	p := []int{}
	if part != nil {
		p = part[0]
		x = arrAdd(p)
	}

	if x == target {
		(*allC) = append((*allC), p)
		return 1
	}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		nn := nums[i+1:]
		total = total + subSum(nn, target, allC, append(p, n))
	}
	return total
}

func findMinL(allC [][]int) int {
	m := 50
	for _, x := range allC {
		if len(x) < m {
			m = len(x)
		}
	}
	return m
}

func lengthCount(allC [][]int, m int) int {
	c := 0
	for _, x := range allC {
		if len(x) == m {
			c += 1
		}
	}
	return c
}

func main() {
	target := 150
	fn := "input/infile17"
	var allC [][]int

	nums := fileToArray(fn)
	t := subSum(nums, target, &allC)
	minL := findMinL(allC)
	lCount := lengthCount(allC, minL)

	fmt.Printf("  %v out of %v match the minimum length of  %v\n", lCount, t, minL)
}
