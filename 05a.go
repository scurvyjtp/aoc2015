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

func checkVowel(v byte) bool {
	if v == 97 || v == 101 || v == 105 || v == 111 || v == 117 {
		return true
	} else {
		return false
	}
}

func checkConcurrent(in string) bool {
	if in[0] == in[1] {
		return true
	} else {
		return false
	}
}

func checkDisallowed(in string) bool {
	if in == "ab" || in == "cd" || in == "pq" || in == "xy" {
		return true
	} else {
		return false
	}
}

func parseLine(line string) bool {
	var vowelCount int = 0
	isConcurrent := false
	isDisallowed := false

	for i := 0; i < len(line); i++ {
		if vowelCount < 3 {
			if checkVowel(line[i]) {
				vowelCount += 1
			}
		}
		if i < len(line)-1 {
			if !(isConcurrent) {
				isConcurrent = checkConcurrent(string(line[i : i+2]))
			}
			if !(isDisallowed) {
				isDisallowed = checkDisallowed(string(line[i : i+2]))
			}
		}
	}

	if vowelCount >= 3 && isConcurrent && !(isDisallowed) {
		fmt.Println("Nice!")
		return true
	} else {
		fmt.Println("Naughty.")
		return false
	}

}

/*****
Rules:
  1) has at least three vowels
  2) has at least one set of concurrent letters
  3) can't contain ab, cd, pq, xy
*****/
func main() {
	fn := "input/infile05"
	niceCount := 0

	lines := fileToArray(fn)
	for _, line := range lines {
		if parseLine(line) {
			niceCount += 1
		}
	}

	fmt.Printf("Nice count: %v\n", niceCount)
}
