package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

/* Global match template */
var tSue = map[string]int {
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
}

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func parseFile (fn string) {
    file, err := os.Open(fn)
    check(err)

    i := 0
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        i += 1
        match := checkLine(scanner.Text())
        if match {
            fmt.Printf("isMatch at %v: %v\n", i,match)
        }
    }
    file.Close()
}

func checkLine(line string) bool{
    a := strings.SplitN(line, ":",2)
    b := strings.Split(strings.Replace(a[1]," ","",-1),",")
    for _,c := range b {
        d := strings.Split(c,":")
        e,_ := strconv.Atoi(d[1])
        if e != tSue[d[0]] {
            return false
        }
    }
    return true
}

func main() {
    fn := "input/infile16"
    parseFile(fn)
}
