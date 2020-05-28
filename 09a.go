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

func checkNodes(path []string, key string) bool {
    for _, p := range path {
        if key == p {
            return true
        }
    }
    return false
}

func generatePermutations(l int, nodes []string, paths *[][]string) {
    if l == 1 {
        sCopy := make([]string, len(nodes))
        copy(sCopy, nodes)
        *paths = append(*paths,sCopy)
    } else {
        for i := 0; i < l-1; i++ {
            generatePermutations(l-1, nodes, paths)
            if l%2 == 0 {
                nodes[i], nodes[l-1] = nodes[l-1], nodes[i]
            } else {
                nodes[0], nodes[l-1] = nodes[l-1], nodes[0]
            }
        }
        generatePermutations(l-1,nodes,paths)
    }
}

func getDistance(v []string, paths map[string]map[string]int) int {
    td := 0
    d := 0
    for i := 1; i < len(v); i++ {
        td = paths[v[i-1]][v[i]]
        if td == 0  {
            td = paths[v[i]][v[i-1]]
            if td == 0 {
                fmt.Printf("Couldn't find connection: %v:%v at offeset %v in %v\n",v[i],v[i-1],i,v )
            }
        }
        d += td
    }
    return d
}

func main() {
    fn := "input/infile09"
    allNodes := make(map[string]map[string]int)

    var path []string
    var allPaths[][]string
    var val int
    var offset int
    var dist int
    first := true

    lines := fileToArray(fn)

    for _, line := range lines {
        r := regexp.MustCompile(`^([A-Za-z]+)\sto\s([A-Za-z]+)\s=\s([0-9]+)$`)
        q := r.FindStringSubmatch(line)
        val, _ = strconv.Atoi(q[3])
        if allNodes[q[1]] == nil {
            allNodes[q[1]] = make(map[string]int)
        }
        allNodes[q[1]][q[2]] = val

        if !(checkNodes(path, q[1])) {
            path = append(path, q[1])
        }

        if !(checkNodes(path, q[2])) {
            path = append(path, q[2])
        }
    }

    generatePermutations(len(path),path,&allPaths)

    for n,v := range allPaths {
        d := getDistance(v,allNodes)
        if first {
            dist = d
            offset = n
            first = false
        }
        if d < dist {
            dist = d
            offset = n
        }
    }
    fmt.Printf("Winner: %v\t path %v: %v\n", dist, offset, allPaths[offset])
}

