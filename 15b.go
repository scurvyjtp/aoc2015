package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func parseArray(l []string, ir *map[string]map[string]int, il *[]string) []string {
    in := []string {}

    for n, l := range l {
        a := strings.Split(l,":")

        in = append(in, a[0])
        if (*ir)[a[0]] == nil {
            (*ir)[a[0]] = make(map[string]int)
        }
        b := strings.Split(a[1], ",") 
        for _,c := range b {

            d := strings.Split(strings.Trim(c, " ")," ")
            e,_ := strconv.Atoi(d[1])
            (*ir)[a[0]][d[0]] = e

            if n == 0 {
                (*il) = append((*il),d[0])
            }
        }
    }
    return in
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

func main() {
    fn := "input/infile15"
    ing := make(map[string]map[string]int)
    ilist := []string{}
    target := 100
    ans := 0
    v := "calories"

    lines := fileToArray(fn)
    in := parseArray(lines, &ing, &ilist)

    for i := 0; i < target; i++ {
        for j := 0; j < target; j++ {
            for k := 0; k < target; k++ {
                for l := 0; l < target; l++ {
                    if i + j + k + l == target{
                        if (ing[in[0]][v]* i)+(ing[in[1]][v]*j)+
                           (ing[in[2]][v]* k)+(ing[in[3]][v]*l) == 500 {

                            total := 1
                            for _,s := range ilist {
                                if s != "calories" {
                                    a := (ing[in[0]][s]* i)+(ing[in[1]][s]*j)+
                                         (ing[in[2]][s]* k)+(ing[in[3]][s]* l)
                                    if a < 0 {
                                        a = 0
                                    }
                                    total = total * a
                                }
                            }
                            if total > ans {
                                ans = total
                            }
                        }
                    }
                }
            }
        }
    }
fmt.Printf("%v\n",ans)
}
