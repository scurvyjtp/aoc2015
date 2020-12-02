package main

import (
	"bufio"
	"fmt"
	"os"
)

const G = 100

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printCurr(curr *[G][G]bool) {
	for i := 0; i < G; i++ {
		for j := 0; j < G; j++ {
			if (*curr)[i][j] == true {
				fmt.Printf("%c", 35)
			} else {
				fmt.Printf("%c", 46)
			}
		}
		fmt.Printf("\n")
	}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func evalCurr(curr *[G][G]bool, i int, j int) bool {
	t := 0

	if (i == 0 || i == G-1) && (j == 0 || j == G-1) {
		return true
	}

	if i != 0 {
		if j != 0 {
			t += btoi(curr[i-1][j-1])
		}
		t += btoi(curr[i-1][j])
		if j != G-1 {
			t += btoi(curr[i-1][j+1])
		}
	}

	if j != 0 {
		t += btoi(curr[i][j-1])
	}
	if j != G-1 {
		t += btoi(curr[i][j+1])
	}

	if i != G-1 {
		if j != 0 {
			t += btoi(curr[i+1][j-1])
		}
		t += btoi(curr[i+1][j])
		if j != G-1 {
			t += btoi(curr[i+1][j+1])
		}
	}

	if curr[i][j] {
		if t == 2 || t == 3 {
			return true
		} else {
			return false
		}
	} else {
		if t == 3 {
			return true
		} else {
			return false
		}
	}
	return false
}

func getNext(curr *[G][G]bool, next *[G][G]bool) {
	for i := 0; i < G; i++ {
		for j := 0; j < G; j++ {
			(*next)[i][j] = evalCurr(curr, i, j)
		}
	}
}

func countOn(curr *[G][G]bool) int {
	t := 0
	for i := 0; i < G; i++ {
		for j := 0; j < G; j++ {
			if (*curr)[i][j] {
				t += 1
			}
		}
	}
	return t
}

func fileToArray(fn string, curr *[G][G]bool) {
	i := 0
	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		for j, m := range scanner.Text() {
			if m == 35 { //on
				(*curr)[i][j] = true
			} else if m == 46 { //off
				(*curr)[i][j] = false
			}
		}
		i += 1
	}
	file.Close()
}

func main() {
	fn := "input/infile18"
	var curr [G][G]bool
	var next [G][G]bool
	var i int

	fileToArray(fn, &curr)
	curr[0][0] = true
	curr[0][99] = true
	curr[99][0] = true
	curr[99][99] = true

	for i = 0; i < 100; i++ {
		getNext(&curr, &next)
		curr = next
	}
	fmt.Printf("Total at %v: %v\n", i, countOn(&curr))
}
