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

func findInSlice (p []posit, x int, y int) int{
    for num, n := range p {
        if x == n.x && y == n.y {
            return num
        }
    }
    return -1
}

type posit struct {
    x       int
    y       int
    offset  int
    count   int
}

func parseInstructions (p []posit, d []rune) []posit {
    var curX, curY, i, fCounter int

    //fmt.Printf("x:\t%v  y:\t%v\n",curX,curY)
    //  60 = <
    //  62 = > 
    //  94 = ^
    // 118 = v

    for i =0; i < len(d); i++ {
        if d[i] == 60 {
            curX = curX - 1
        } else if d[i] == 62 {
            curX = curX + 1
        } else if d[i] == 94 {
            curY = curY + 1
        } else if d[i] == 118 {
            curY = curY - 1
        } else {
            fmt.Println("Invalid value.\n")
            os.Exit(1)
        }


        new_p := posit{
            x: curX,
            y: curY,
            offset: i,
            count: 1,
        }

        j := findInSlice(p,curX,curY)
        if j == -1 {
            p = append(p,new_p)
        } else {
            fCounter += 1
            p[j].count = p[j].count + 1;
        }
    }
    return p
}


func main() {
    var fn = "input/infile03"
    var houses []posit
    var santa, robot []rune

    p := posit{
        x: 0,
        y: 0,
        offset: 0,
        count: 1,
    }

    houses = append(houses, p)
    lines := fileToArray(fn)

    for _, line := range lines {
        for n, c := range line {
            if n%2 == 0 {
                santa = append(santa,c)
            } else {
                robot = append(robot,c)
            }
        }
    }

    houses = parseInstructions(houses, santa)
    houses = parseInstructions(houses, robot)

    fmt.Printf("Houses Visited: %v\n", len(houses))

}
