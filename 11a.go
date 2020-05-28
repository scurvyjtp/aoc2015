package main

import (
    "fmt"
    "os"
)

func checkIJO(s string) bool {
    for _,i := range s {
        if i == 'i' || i == 'j' || i == 'o' {
            return true
        }
    }
    return false
}

func checkConsecutive(s string) bool {
    matchCount := 0
    var matchChar byte

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            if matchCount == 1 && s[i] != matchChar {
                return true
            }
            matchCount = 1
            matchChar = s[i]
            i += 1      // skip a char to non-overlap
        }
    }
    return false
}

func checkIncremental(s string) bool {
    for i :=0; i< len(s)-2; i++ {
        if s[i+1] == s[i]+1 && s[i+2] == s[i+1]+1 && s[i+2] <= 122 {
            return true
        }
    }
    return false
}



func isValid (s string) bool {
    if !(checkIJO(s)) && checkConsecutive(s) && checkIncremental(s) {
        return true
    }
    return false
}

func bitMath(s string, o int) string{
    // a:  97
    // z: 122
    sb := []byte(s)
    sb[o] = sb[o] + 1
    if sb[o] == 123 {
        sb[o] = 97 
        s = bitMath(string(sb), o-1)
    } else {
        s = string(sb)
    }
    return s
}


func main() {
    input := "vzbxkghb" // 10a input
    input =  "vzbxxyzz" // 10b input
    n := 0

    for true {
        n += 1
        if n > 100000000 {
            os.Exit(5)
        }
        input = bitMath(input,len(input)-1)
        if(isValid(input)) {
            fmt.Printf("Answer: %v:\t%v\n", n,input)
            break
        }
    }


}
