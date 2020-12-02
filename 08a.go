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
		//fmt.Printf("%v\tl:\t%v\tm:\t%v\tEC:\t%v\tt|f:\t%v\n",string(line[i]),l,m,escapeCount,(line[i]=='\\' && escapeCount ==0))
	}
	counts["l"] = counts["l"] + l
	counts["m"] = counts["m"] + m
	fmt.Printf("%v\n\t%v - %v = %v\n", line, l, m, (l - m))
}

func main() {
	fn := "input/infile08"
	counts := map[string]int{
		"l": 0, //literals
		"m": 0, //memory
	}

	lines := fileToArray(fn)
	for _, line := range lines {
		parseLine(line, counts)
		//fmt.Println(line)
	}

	fmt.Printf("%v - %v = %v\n", counts["l"], counts["m"], counts["l"]-counts["m"])
}
