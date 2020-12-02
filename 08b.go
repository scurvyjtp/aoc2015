package main

import (
	"bufio"
	"fmt"
	"os"
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

func parseLine(line string, counts map[string]int) {
	l := len(line)
	m := 0
	escapeCount := 0

	for i := 0; i < len(line); i++ {
		if escapeCount > 0 {
			escapeCount -= 1
		} else if escapeCount == 0 {
			if line[i] == '\\' {
				if line[i+1] == '\\' || line[i+1] == '"' {
					m += 1
					escapeCount = 1
				} else if line[i+1] == 'x' {
					m += 1
					escapeCount = 3
				}
			} else {
				if !(i == 0 || i == len(line)-1) {
					m += 1
				}
			}
		} else {
			fmt.Printf("Error: We should never get here!")
			os.Exit(1)
		}
	}
	counts["l"] = counts["l"] + l
	counts["m"] = counts["m"] + m
}

func encodeLine(line string) string {
	var newL string
	newL += string(34)
	for _, l := range line {
		if string(l) == string(34) || string(l) == string(92) {
			newL += string(92)
		}
		newL += string(l)
	}
	newL += string(34)
	return newL
}

func main() {
	fn := "input/infile08"
	counts := map[string]int{
		"l": 0, //literals
		"m": 0, //memory
	}
	newCounts := map[string]int{
		"l": 0, //literals
		"m": 0, //memory
	}

	var newLine []string

	lines := fileToArray(fn)
	for _, line := range lines {
		newLine = append(newLine, encodeLine(line))
		parseLine(line, counts)
	}

	for _, line := range newLine {
		parseLine(line, newCounts)
	}

	fmt.Printf("New Total Chars - Old Memory Footprint:\n\t%v - %v = %v\n", newCounts["l"], counts["l"], newCounts["l"]-counts["l"])
}
