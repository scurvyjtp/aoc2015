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

func parseInstruction(k string, n map[string]string, v map[string]uint16) uint16 {
	if v[k] != 0 {
		return v[k]
	}

	r := regexp.MustCompile(`^.*(NOT|AND|OR|RSHIFT|LSHIFT).*$`)
	s := r.FindStringSubmatch(n[k])
	if s != nil {
		r := regexp.MustCompile(`([a-z]+|[0-9+])?\s?(NOT|AND|OR|RSHIFT|LSHIFT)\s([a-z]+|[0-9]+)`)
		s := r.FindStringSubmatch(n[k])

		switch s[2] {
		case "NOT":
			v[k] = ^(parseInstruction(s[3], n, v))
			return v[k]
		case "AND":
			t, e := strconv.Atoi(s[1])
			if e == nil {
				v[k] = uint16(t) & parseInstruction(s[3], n, v)
			} else {
				v[k] = parseInstruction(s[1], n, v) & parseInstruction(s[3], n, v)
			}
			return v[k]
		case "OR":
			v[k] = parseInstruction(s[1], n, v) | parseInstruction(s[3], n, v)
			return v[k]
		case "RSHIFT":
			t, _ := strconv.Atoi(s[3])
			v[k] = parseInstruction(s[1], n, v) >> t
			return v[k]
		case "LSHIFT":
			t, _ := strconv.Atoi(s[3])
			v[k] = parseInstruction(s[1], n, v) << t
			return v[k]
		}
		fmt.Printf("%v\n", s[1])
	} else {
		/* If there is no instruction it's a root node
		 * Convert the response to an integer
		 * If it's not an integer, it's a register
		 * return the result of parseInstruction
		 */
		t, e := strconv.Atoi(n[k])
		if e == nil {
			v[k] = uint16(t)
			return v[k]
		} else {
			v[k] = parseInstruction(n[k], n, v)
			return v[k]
		}
		fmt.Printf("No instruction! -- %v\n", n[k])
	}

	fmt.Printf("We should never be here!\n")
	return 0
}

func main() {
	fn := "input/infile07"
	nodes := make(map[string]string)
	values := make(map[string]uint16)

	lines := fileToArray(fn)
	for _, line := range lines {
		r := regexp.MustCompile(`(.*)\s->\s([a-z]+)$`)
		s := r.FindStringSubmatch(line)
		nodes[s[2]] = s[1]
	}

	values["b"] = 46065 // Answer from 07a
	key := "a"
	_ = parseInstruction(key, nodes, values)
	fmt.Printf("%v: %v\n", key, values[key])
}
