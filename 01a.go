package main

import (
    "fmt"
    "io/ioutil"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var up int = 0
    var down int = 0

    data, err := ioutil.ReadFile("input/infile01")
    check(err)

    for i := 0; i < len(data); i++ {
        if data[i] == 40 {
            up += 1
        } else if data[i] == 41 {
            down += 1
        }
    }

    fmt.Println((up - down))
}

