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

func fileToArray(fn string, rules *map[string]string) string {
	file, err := os.Open(fn)
	check(err)
	var ans string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m, _ := regexp.MatchString("=>", scanner.Text())
		if m {
			a := strings.Split(scanner.Text(), " => ")
			(*rules)[reverseString(a[1])] = reverseString(a[0])
		} else {
			if len(scanner.Text()) > 0 {
				ans = reverseString(scanner.Text())
			}
		}
	}
	file.Close()
	return ans
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

func profileM(m string) {
	var j string
	tCnt := 0
	elCnt := make(map[string]int)

	for i := 0; i < len(m); i++ {
		if m[i+1] >= 97 && m[i+1] <= 122 {
			j = string(m[i : i+2])
			i += 1
		} else {
			j = string(m[i])
		}
		elCnt[j] += 1
		tCnt += 1
	}
	fmt.Printf("Total Elements: %v\n", tCnt)
	//    for k,v := range elCnt {
	//        fmt.Printf("%v:%v\n",k,v)
	//    }
	fmt.Printf("Format:\n\t Tot - (Rn + Ar) - (2 * Y) - 1 = Ans\n")
	fmt.Printf("%v - (%v + %v) - (2 * %v) - 1 = %v\n", tCnt, elCnt["Rn"], elCnt["Ar"], elCnt["Y"], tCnt-(elCnt["Rn"]+elCnt["Ar"])-(2*elCnt["Y"])-1)
}

func printRules(rules *map[string]string) {
	for k, v := range *rules {
		fmt.Printf("%v => %v\n", k, v)
	}
}
func reverseString(in string) string {
	rs := []rune(in)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	fn := "input/infile19"
	rules := make(map[string]string)
	m := fileToArray(fn, &rules)

	fmt.Printf("%v\n", m)

	//    printRules(&rules)
	//    profileM(m)
}
