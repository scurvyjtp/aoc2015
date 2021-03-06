package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	in := "yzbqklnj"
	i := 0

	for {
		t := in + strconv.Itoa(i)
		m := md5.Sum([]byte(t))
		s := hex.EncodeToString(m[:])

		if s[0:6] == "000000" {
			fmt.Printf("Success at %v: %v\n", i, s)
			os.Exit(0)
		}
		i += 1
	}

}
