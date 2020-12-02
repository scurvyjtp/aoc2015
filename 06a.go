package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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

type command struct {
	inst   string
	startX int64
	startY int64
	stopX  int64
	stopY  int64
}

func parseLine(l string) command {
	var s command
	r := regexp.MustCompile(`.*(?P<inst>on|off|toggle)\s(?P<startX>\d+),(?P<startY>\d+)\sthrough\s(?P<stopX>\d+),(?P<stopy>\d+)`)
	q := r.FindStringSubmatch(l)
	s.inst = q[1]
	s.startX, _ = strconv.ParseInt(q[2], 10, 64)
	s.startY, _ = strconv.ParseInt(q[3], 10, 64)
	s.stopX, _ = strconv.ParseInt(q[4], 10, 64)
	s.stopY, _ = strconv.ParseInt(q[5], 10, 64)
	return s
}

func parseInstruction(c command, l *[1000][1000]bool) {
	for i := c.startX; i <= c.stopX; i++ {
		for j := c.startY; j <= c.stopY; j++ {
			if c.inst == "off" {
				l[i][j] = false
			} else if c.inst == "on" {
				l[i][j] = true
			} else if c.inst == "toggle" {
				l[i][j] = !(l[i][j])
			} else {
				fmt.Printf("Not a valid instruction! %v", c.inst)
			}
		}
	}

}

func main() {
	fn := "input/infile06"
	var c []command
	var lights [1000][1000]bool
	var lightCount int

	lines := fileToArray(fn)
	for _, line := range lines {
		c = append(c, parseLine(line))
	}

	for i := 0; i < len(c); i++ {
		parseInstruction(c[i], &lights)
	}

	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if lights[i][j] {
				lightCount += 1
			}
		}
	}

	fmt.Printf("Lights that are on: %v\n", lightCount)
}
