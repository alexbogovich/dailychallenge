package main

func IsArrayContainsGivenSum(arr []int, sum int) bool {
    encounteredNumbers := make(map[int]bool)
    for _, v := range arr {
        diff := sum - v
        if encounteredNumbers[diff] {
           return true
        }
        encounteredNumbers[v] = true
    }
    return false
}
