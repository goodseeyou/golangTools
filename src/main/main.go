package main

import (
    "list"
    "fmt"
    "math/rand"
    "time"
    "sort"
)

const (
    iterTimes = 10000000 << 1
    numberRange = 1 << 24
    k = iterTimes/3
)

func main() {
    rand.Seed(time.Now().UnixNano())

    var randomInts = make([]int, 0, iterTimes)
    for i := 0; i< iterTimes ; i++ {
        randomInts = append(randomInts, rand.Intn(numberRange))
    }

    fmt.Println(time.Now())
    selected, _ := list.SelectTheKthSmallNumber(randomInts, k)
    fmt.Println(selected)

    fmt.Println(time.Now())

    sort.Ints(randomInts)
    fmt.Println(randomInts[k-1])
    fmt.Println(time.Now())

    if randomInts[k-1] != selected {
        fmt.Println(randomInts[k-4:k+3])
    }
}
