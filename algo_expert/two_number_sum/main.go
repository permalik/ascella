package main

import (
    "log"
)

func p(msg string, val any) {

    log.Printf("msg::%s\n%v", msg, val)
}

func twoNumberSum(a []int, b int) []int {

    var c []int
    for i := 0; i < len(a); i++ {
        p("i::\n", i)
        numOne := a[i]
        ii := i + 1
        numTwo := a[ii]
        test := []int{numOne, numTwo}
        p("test::\n", test)

        if numOne + numTwo == b {
            c = append(c, numOne, numTwo)
            p("c::\n", c)
            return c
        }
    }

    return c
}

func main() {

    arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
    sum := 10
    twoNumberSum(arr, sum)
}

