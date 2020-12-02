package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func min3(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}

func surfaceArea(l, w, h int) int {
	return ((2 * l * w) + (2 * w * h) + (2 * h * l))
}

func main() {
	var lines []string
	var m []string
	var area, total int
	var la, ha, wa, ea int
	var l, w, h int

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

		l, _ = strconv.Atoi(m[0])
		w, _ = strconv.Atoi(m[1])
		h, _ = strconv.Atoi(m[2])

		la = l * w
		wa = w * h
		ha = h * l
		ea = min3(la, wa, ha)

		area = 2*la + 2*wa + 2*ha + ea

		total = area + total
		//fmt.Println(lineNum, line, la, wa, ha, ea, area, total)

	}
	fmt.Println("Total:", total)
}
