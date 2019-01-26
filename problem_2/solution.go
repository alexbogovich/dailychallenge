package main

// a[i] = a[0] * ... * a[i-2] * a[i-1] * a[i+1] * a[i+2] * ... * a[n]
// a[i] = a[0] * ... * a[i-2] * a[i-1] * GetMultiplyFromBack(a)[i]
func ProductionExceptIndex(arr []int) []int {
	length := len(arr)
	res := make([]int, length, length)
	iToNMul := GetMultiplyFromBack(arr)
	memo := 1

	for i, v := range arr {
		res[i] = memo * iToNMul[i]
		memo *= v
	}

	return res
}

// a[i] = a[i+1] * a[i+2] * ... * a[n]
func GetMultiplyFromBack(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return []int{0}
	}
	production := make([]int, length, length)
	memo := 1
	for i := length - 1; i >= 0; i-- {
		production[i] = memo
		memo *= arr[i]
	}
	return production
}
