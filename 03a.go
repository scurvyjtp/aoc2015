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

func fileToArray (fn string) []string {
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

func findInSlice (p []posit, x int, y int) bool{

    for _, n := range p {
        if x == n.x && y == n.y {
            return true
        }
    }
    return false
}

type posit struct {
    x       int
    y       int
    offset  int
    count   int
}

func main() {
    var fn = "input/infile03"
    var i,j int
    var positions []posit
    var curX, curY int

    p := posit{
        x: 0,
        y: 0,
        offset: 0,
        count: 1,
    }
    positions = append(positions,p)

    lines := fileToArray(fn)

    for _, line := range lines {
        for i = 0; i < len(line); i++ {
            //  60 = <
            //  62 = > 
            //  94 = ^
            // 118 = v
            //fmt.Printf("%c\t%v\n", line[i], line[i])
            if line[i] == 60 {
                curX = curX - 1
            } else if line[i] == 62 {
                curX = curX + 1
            } else if line[i] == 94 {
                curY = curY + 1
            } else if line[i] == 118 {
                curY = curY - 1
            } else {
                fmt.Println("Invalid value.\n")
                os.Exit(1)
            }

            p := posit{
                x: curX,
                y: curY,
                offset: i,
                count: 1,
            }
            if findInSlice(positions,curX,curY) {
                //fmt.Println("Found!")
                positions[j].count = positions[j].count + 1;
            } else {
                j += 1
                positions = append(positions,p)
            }

            //fmt.Printf("Current Position: %v,%v.  Strucht: %v\n", curX, curY, positions[i])
        }
    }
    fmt.Printf("Positions Count: %v\n", len(positions))
}
