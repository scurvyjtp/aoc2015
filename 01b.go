package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var tally int = 0

	data, err := ioutil.ReadFile("input/infile01")
	check(err)

	for i := 0; i < len(data); i++ {
		if data[i] == 40 {
			tally += 1
		} else if data[i] == 41 {
			tally -= 1
		}
		if tally == -1 {
			fmt.Printf("Basement at position: %v\n", i+1)
			os.Exit(0)
		}

	}

}
