package main

// O(2n)
// mem n
// additional mem n + 1
func WithMap(arr []int) int {
	existingNums := make(map[int]bool)
	max := 0
	for _, v := range arr {
		existingNums[v] = true
		if max < v {
			max = v
		}
	}

	for i := 1; i < max; i++ {
		if !existingNums[i] {
			return i
		}
	}

	return max + 1
}
