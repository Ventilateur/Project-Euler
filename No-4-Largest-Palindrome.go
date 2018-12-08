package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)


func findFactor(in uint64) bool {
    var f uint64 = 10
    for ;f <= 90; f++ {
        if (in % f == 0) && (in / f >= 100) && (in / f <= 999) {
            return true
        }
    }
    return false
}


func closestPalindrome(in uint64) uint64 {
    n := in - 1
    a := uint64(n / 100001)
    b := uint64((n - a * 100001) / 10010)
    c := uint64((n - a * 100001 - b * 10010) / 1100)
    return a * 100001 + b * 10010 + c *1100
}



func compute(input []uint64) []uint64 {
    var output []uint64
    for _, N := range input {
        ghk := uint64(closestPalindrome(N) / 1000)
        var elevenF uint64
        for ghk >= 100 {
            elevenF = 9091 * uint64(ghk / 100) + 910 * ((ghk / 10) % 10) + 100 * (ghk % 10)
            if findFactor(elevenF) {
                output = append(output, elevenF * 11)
                break
            }
            ghk--
        }
    }
    return output
}

func main() {
    var input []uint64
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        aInt, _ := strconv.ParseUint(scanner.Text(), 10, 64)
        input = append(input, aInt)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    
    for _, v := range compute(input[1:]) {
        fmt.Println(v)
    }
}
