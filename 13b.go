package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "strconv"
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

func inArray (s string, a []string) bool {
    for _,v := range a {
        if v == s {
            return true
        }
    }
    return false
}

func permuteAttendees(n int, a []string, seating *[][]string) {
    if n == 1 {
        s := make([]string, len(a))
        copy(s, a)
        *seating = append(*seating,s)
    } else {
        permuteAttendees(n-1, a, seating)

        for i := 0; i < n-1; i++ {
            if n % 2 == 0 {
                a[i], a[n-1] = a[n-1], a[i]
            } else {
                a[0], a[n-1] = a[n-1], a[0]
            }
            permuteAttendees(n-1,a, seating)
        }
    }
}

func scoreArray(a []string, happyMap *map[string]map[string]int) int {
    total := 0

    for i := 0; i < len(a); i++ {
        if i == 0 {
            total += (*happyMap)[a[i]][a[len(a)-1]] + (*happyMap)[a[i]][a[i+1]]
        } else if i == len(a)-1 {
            total += (*happyMap)[a[i]][a[i-1]] + (*happyMap)[a[i]][a[0]]
        } else {
            total += (*happyMap)[a[i]][a[i-1]] + (*happyMap)[a[i]][a[i+1]]
        }
    }
    return total
}

func maxArray (a []int) int {
    r := 0
    for _,v := range a {
        if v > r {
            r = v
        }
    }
    return r
}

func main() {
    fn := "input/infile13"
    attendees := []string{}
    seating := [][]string{}
    scores := []int{}
    happyMap := make(map[string]map[string]int)

    r := regexp.MustCompile(`^(.*)\swould\s(gain|lose)\s(\d+).*\s(.*)\.$`)

    lines := fileToArray(fn)
    for _, line := range lines {
        q := r.FindStringSubmatch(line)
        if !(inArray(q[1], attendees)) {
            attendees = append(attendees, q[1])
        }

        if happyMap[q[1]] == nil {
            happyMap[q[1]] = make (map[string]int)
        }

        v,_ := strconv.Atoi(q[3])
        if q[2] == "lose" {
            happyMap[q[1]][q[4]] = -(v)
        } else {
            happyMap[q[1]][q[4]] = v
        }

    }

    happyMap["Me"] = make(map[string]int)
    for _,v := range attendees {
        happyMap["Me"][v] = 0
    }
    attendees = append(attendees, "Me")

    permuteAttendees(len(attendees), attendees, &seating)

    for _,v := range seating {
        t := scoreArray(v, &happyMap)
        scores = append(scores,t)
    }

    s := maxArray(scores)
    fmt.Printf("%v\n",s)
}
