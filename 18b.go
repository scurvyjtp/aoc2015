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

func printStamp(curr *[100][100]bool) {
    fmt.Printf("%v%v\n%v%v\n\n",curr[0][0],curr[0][99],curr[99][0],curr[99][99])
}

func printCurr(curr *[100][100]bool) {
    for i := 0; i < 100; i++ {
        fmt.Printf("%v: ",i)
        for j := 0; j < 100; j++ {
            if (*curr)[i][j] == true {
                fmt.Printf("%c",35)
            } else {
                fmt.Printf("%c",46)
            }
        }
        fmt.Printf("\n")
    }
}

func btoi (b bool) int {
    if b {
        return 1
    }
    return 0
}

func evalCurr(curr *[100][100]bool, i int, j int) bool{
    t := 0

    if (i == 0 || i == 99) && (j == 0 ||  j == 99) {
        return true
    }

    if i != 0 {
        if j != 0  { t += btoi(curr[i-1][j-1]) }
        t += btoi(curr[i-1][j])
        if j != 99 { t += btoi(curr[i-1][j+1]) }
    }

    if j != 0  { t += btoi(curr[i][j-1]) }
    if j != 99 { t += btoi(curr[i][j+1]) }

    if i != 99 {
        if j != 0  { t += btoi(curr[i+1][j-1]) }
        t += btoi(curr[i+1][j])
        if j != 99 { t += btoi(curr[i+1][j+1]) }
    }

    if curr[i][j] {
        if (t == 2 || t == 3) {
            return true
        } else {
            return false
        }
    } else {
        if t == 3 {
            return true
        } else {
            return false
        }
    }
    return false
}

func getNext(curr *[100][100]bool, next *[100][100]bool) {
    for i := 0; i < 100; i++{
        for j := 0; j < 100; j++{
            (*next)[i][j] = evalCurr(curr, i, j)
            //fmt.Printf("%v\n", evalCurr(curr, i, j));
        }
    }
}

func countOn (curr *[100][100]bool) int {
    t := 0
    for i := 0; i < 100; i++{
        for j := 0; j < 100; j++{
            if (*curr)[i][j] {
                t += 1
            }
        }
    }
    return t
}

func fileToArray (fn string, curr *[100][100]bool) {
    i := 0
    file, err := os.Open(fn)
    check(err)

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        for j,m := range scanner.Text() {
            if m == 35 { //on
                (*curr)[i][j] = true
            } else if m == 46 { //off
                (*curr)[i][j] = false
            }
        }
        i += 1
    }
    file.Close()
}

func main() {
    fn := "input/infile18"
    var curr [100][100]bool
    var next [100][100]bool
    var i int

    fileToArray(fn, &curr)
    for i = 0; i < 100; i++ {
        //printCurr(&curr)
        //printStamp(&curr)
        getNext(&curr, &next)
        curr = next
    }
        fmt.Printf("Total at %v: %v\n",i, countOn(&curr))
}

