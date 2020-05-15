package main

import (
    "fmt"
    "os"
    "bufio"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func fileToArray (fn string) []string {
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

func checkConcurrent (in string) bool {
    if in[0] == in[1] {
        return true
    } else {
        return false
    }
}

func findSubstr (val string, search string) bool {
	slen := len(val)
	for i := 0; i < len(search) - (slen-1); i++ {
		if val == string(search[i:i+2]) {
			return true
		}
	}
	return false
}

func parseLine(line string) bool{
    isDuped:= false
	isMatch := false

    for i := 0; i < len(line); i++ {
        if !(isDuped) && i < len(line) - 3 {
			isDuped = findSubstr(string(line[i:i+2]), string(line[i+2:]))
        }

		if !(isMatch) && i < len(line) -2 {
			isMatch = (line[i] == line[i+2])
		}
    }

	if isMatch && isDuped {
		return true
	} else {
		return false
	}
}

func main() {
    fn := "input/infile05"
    niceCount := 0

    lines := fileToArray(fn)
    for _, line := range lines {
        if(parseLine(line)) {
            niceCount += 1
        }
    }

    fmt.Printf("Nice count: %v\n", niceCount)
}

