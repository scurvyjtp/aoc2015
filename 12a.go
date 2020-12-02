package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []string {
	var l []string
	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}

	file.Close()
	return l
}

func main() {
	fn := "input/infile12"
	total := 0

	r := regexp.MustCompile(`-?[0-9]+`)

	lines := fileToArray(fn)

	for _, line := range lines {
		total = 0
		s := r.FindAllString(line, -1)

		//fmt.Printf("%v:", line)

		for _, t := range s {
			u, _ := strconv.Atoi(t)
			total += u
		}

		fmt.Printf("\t Total: %v\n", total)
	}
}
