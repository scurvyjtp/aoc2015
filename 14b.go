package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "strconv"
)

type deer struct {
    speed int
    time int
    rest int
}


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

func parseArray(l []string, rd *map[string]deer) {

    r := regexp.MustCompile(`^(.*)\scan.*\s(\d+)\skm/s\sfor\s(\d+)\s.*\s(\d+)\sseconds.$`)
    for _, l := range l {
        q := r.FindStringSubmatch(l)
        spd,_ := strconv.Atoi(q[2])
        tim,_ := strconv.Atoi(q[3])
        rst,_ := strconv.Atoi(q[4])
        (*rd)[q[1]] = deer{
            speed: spd,
            time:  tim,
            rest:  rst,
        }
    }
}

func raceDistance(rl int, speed int, time int, rest int) int {
    lastLeg := 0
    if (
        rl % (time+ rest) > time) {
        lastLeg = time * speed
    } else {
        lastLeg = (rl % (time +rest)) * speed
    }
    return (rl / (time + rest) * (time*speed)) + lastLeg
}

func findLeader(dist int, rd *map[string]deer, lb *map[string]int) {
    temp := make(map[string]int)
    var leaders []string

    for k,v := range *rd {
        temp[k] = raceDistance(dist,v.speed,v.time,v.rest)
    }

    d := 0
    for k,v := range temp {
        if v > d {
            d = v
            leaders = []string{k}
        } else if v == d {
            leaders = append(leaders,k)
        }
    }

    for _,k := range leaders {
        (*lb)[k] = (*lb)[k] + 1
    }
}

func initLb(rd map[string]deer, lb *map[string]int) {
    for k,_ := range rd {
        (*lb)[k] = 0
    }

}

func printLb(lb map[string]int) {
    for k,v := range lb{
        fmt.Printf("%v: %v\n",k,v)
    }
}

func main() {
    fn := "input/infile14"
    rl := 2503
    rd := make(map[string]deer)
    lb:= make (map[string]int)

    lines := fileToArray(fn)
    parseArray(lines, &rd)

    initLb(rd, &lb)

    for i:=1; i<=rl; i++ {
        findLeader(i,&rd,&lb)
    }

    printLb(lb)
}
