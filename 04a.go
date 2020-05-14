package main

import (
    "fmt"
    "crypto/md5"
    "os"
    "strconv"
    "encoding/hex"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    in := "yzbqklnj"
    i := 0

    for {
        t := in + strconv.Itoa(i)
        m := md5.Sum([]byte(t))
        s := hex.EncodeToString(m[:])

        if s[0:5] == "00000"  {
            fmt.Printf("Success at %v: %v\n", i,s)
            os.Exit(0)
        }
        i += 1
    }

}

