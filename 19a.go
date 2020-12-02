package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string, rules *map[string][]string) string {
	file, err := os.Open(fn)
	check(err)
	var in string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m, _ := regexp.MatchString("=>", scanner.Text())
		if m {
			a := strings.Split(scanner.Text(), " => ")
			if (*rules)[a[0]] == nil {
				(*rules)[a[0]] = make([]string, 0)
			}
			(*rules)[a[0]] = append((*rules)[a[0]], a[1])

		} else {
			if len(scanner.Text()) > 0 {
				in = scanner.Text()
			}
		}
	}
	file.Close()
	return in
}

func inToArr(in string) []string {
	var out []string
	var j string

	for i := 0; i < len(in); i++ {
		if in[i+1] >= 97 && in[i+1] <= 122 {
			j = string(in[i : i+2])
			i += 1
		} else {
			j = string(in[i])
		}
		out = append(out, j)
	}
	return out
}

func popArr(rules map[string][]string, in string, all *[]string) {
	for k, v := range rules {
		for _, tr := range v {
			r := regexp.MustCompile(k)
			offsets := r.FindAllStringIndex(in, -1)

			for _, o := range offsets {
				out := in[0:o[0]] + tr + in[o[1]:]
				if !(checkArr(all, out)) {
					(*all) = append((*all), out)
				}
			}
		}
	}
}

func checkArr(all *[]string, in string) bool {
	for _, v := range *all {
		if in == v {
			return true
		}
	}
	return false
}

func main() {
	fn := "input/infile19"
	rules := make(map[string][]string)
	var all []string
	in := fileToArray(fn, &rules)

	popArr(rules, in, &all)
	fmt.Printf("Raw Subs: %v\n", len(all))
}
