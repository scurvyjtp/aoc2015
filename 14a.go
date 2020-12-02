package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type deer struct {
	speed int
	time  int
	rest  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []string {
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
		spd, _ := strconv.Atoi(q[2])
		tim, _ := strconv.Atoi(q[3])
		rst, _ := strconv.Atoi(q[4])
		(*rd)[q[1]] = deer{
			speed: spd,
			time:  tim,
			rest:  rst,
		}
	}
}

func raceDistance(rl int, speed int, time int, rest int) int {
	lastLeg := 0
	if rl%(time+rest) > time {
		lastLeg = time * speed
	} else {
		lastLeg = (rl % (time + rest)) * speed
	}
	return (rl / (time + rest) * (time * speed)) + lastLeg
}

func main() {
	fn := "input/infile14"
	rl := 2503
	rd := make(map[string]deer)

	lines := fileToArray(fn)
	parseArray(lines, &rd)

	for k, v := range rd {
		fmt.Printf("  %v:\t%v\n", k, raceDistance(rl, v.speed, v.time, v.rest))
	}
}
