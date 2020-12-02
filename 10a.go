package main

import (
	"fmt"
	"strconv"
)

func encode(in []int) []int {
	out := make([]int, 1)
	count := 1
	value := in[0]
	first := true

	for i := 1; i < len(in); i++ {
		if in[i] == in[i-1] { // if values are the ame increment count
			count += 1
		} else { // if values differ reset value and count
			if first {
				first = false
				out[0] = count
				out = append(out, value)
			} else {
				out = append(out, count)
				out = append(out, value)
			}
			value = in[i]
			count = 1
		}
	}
	out = append(out, count)
	out = append(out, value)
	return out
}

func main() {
	in := "3113322113"
	out := make([][]int, 41)
	var tmp []int

	for i := 0; i < len(in); i++ {
		v, _ := strconv.ParseInt(string(in[i]), 10, 64)
		tmp = append(tmp, int(v))
	}
	out[0] = tmp

	for i := 1; i < 41; i++ {
		x := encode(out[i-1])
		out[i] = x
	}

	fmt.Printf("length: %v\n", len(out[(len(out)-1)]))
}
