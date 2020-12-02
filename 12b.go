package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkJ(j interface{}) int {
	n := 0

	switch j := j.(type) {

	case []interface{}:
		for _, v := range j {
			n += checkJ(v)
		}
	case map[string]interface{}:
		isRed := false

		for _, v := range j {
			s, ok := v.(string)
			if ok && s == "red" {
				isRed = true
				break
			}
		}

		if !(isRed) {
			for _, v := range j {
				n += checkJ(v)
			}
		}

	case float64:
		n = int(j)
	}

	return n
}

func main() {
	fn := "input/infile12"
	var j map[string]interface{}

	file, err := ioutil.ReadFile(fn)
	check(err)

	e := json.Unmarshal([]byte(file), &j)
	check(e)

	fmt.Printf("%v\n", checkJ(j))
}
