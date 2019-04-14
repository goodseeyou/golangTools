package list

import (
    "fmt"
    "sort"
)


func SelectTheKthSmallNumber(list []int, k int) (int, error) {
    const numberOfEachColumn = 5

    switch {
        case list == nil:
            return 0, fmt.Errorf("invalid parameter: list should not be nil")
        case len(list) < k:
            return 0, fmt.Errorf("invalid parameter: k (%d) should smaller than length of list (%d)", k, len(list))
        case k <= 0:
            return 0, fmt.Errorf("invalid parameter: k (%d) should be positive integer. list (%d)", k, len(list))
    }

    if len(list) <= numberOfEachColumn {
        sort.Ints(list)
        return list[k-1], nil
    }

    var sortedColumnTable = generateSortedColumnTable(list, numberOfEachColumn)
    var medianList = getMedianList(sortedColumnTable)

    kOfMedian := len(medianList) / 2
    if kOfMedian == 0 {
        kOfMedian = 1
    }
    medianOfMedian, err := SelectTheKthSmallNumber(medianList, kOfMedian)
    if err != nil {
        return 0, fmt.Errorf("error while get median of median: %s", err)
    }

    smallerList, equalList, biggerList := distributeList(sortedColumnTable, medianOfMedian)

    switch {
        case k <= len(smallerList):
            return SelectTheKthSmallNumber(smallerList, k)

        case k > len(smallerList) + len(equalList):
            return SelectTheKthSmallNumber(biggerList, k - (len(smallerList) + len(equalList)))

        default: // smallerList < k < biggerList, in other words, k locate in equalList
            return equalList[0], nil
    }

}


func generateSortedColumnTable(list []int, numberOfEachColumn int) ([][]int) {
    var medianListSize = len(list)/numberOfEachColumn
    if len(list) % numberOfEachColumn > 0 {
        medianListSize++
    }

    var table [][]int = make([][]int, 0, medianListSize)
    for i := 0; i < len(list); i+=numberOfEachColumn {
        var column []int
        if i+numberOfEachColumn > len(list) {
            column = list[i:]
        } else {
            column = list[i:i+numberOfEachColumn]
        }
        sort.Ints(column)
        table = append(table, column)
    }

    return table
}


func getMedianList(sortedColumnTable [][]int) []int {
    var medianList []int = make([]int, 0, len(sortedColumnTable))
    for _, column := range sortedColumnTable {
        median := column[len(column) / 2]
        medianList = append(medianList, median)
    }

    return medianList
}


func distributeList(sortedColumnTable [][]int, medianOfMedian int) ([]int, []int, []int) {
    if len(sortedColumnTable) <= 0 {
        return []int{}, []int{}, []int{}
    }
    approximateSizeOfElementInTable := len(sortedColumnTable) * len(sortedColumnTable[0])

    var smallerList, equalList, biggerList []int = make([]int, 0, approximateSizeOfElementInTable / 4), make([]int, 0), make([]int, 0, approximateSizeOfElementInTable / 4)
    for _, column := range sortedColumnTable {
        for _, v := range column {
            switch {
                case v < medianOfMedian:
                    smallerList = append(smallerList, v)
                case v == medianOfMedian:
                    equalList = append(equalList, v)
                case v > medianOfMedian:
                    biggerList = append(biggerList, v)
            }
        }
    }

    return smallerList, equalList, biggerList
}

