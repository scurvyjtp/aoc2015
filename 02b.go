package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var lines, m []string
	n := []int{0, 0, 0}
	var i, total int

	file, err := os.Open("input/infile02")
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	for _, line := range lines {
		m = strings.Split(line, "x")

		for i = 0; i < len(m); i++ {
			n[i], _ = strconv.Atoi(m[i])
		}
		sort.Ints(n)
		total = total + ((2 * n[0]) + (2 * n[1]) + (n[0] * n[1] * n[2]))
	}

	fmt.Println("Total:", total)
}
