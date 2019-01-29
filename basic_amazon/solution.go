package main

import "fmt"

func fibonacci(n int) int {
	x, y, i := 0, 1, 0
	for i < n {
		x, y = x+y, x
		i++
	}
	return x
}

func factorial(n int) int {
	if n >= 1 {
		return n * factorial(n-1)
	}
	return 1
}

// c = m! / (((m - 2)!) * 2!)
func C2(m int) int {
	return factorial(m) / (factorial(m-2) * factorial(2))
}

type IntMap = map[int]bool

func CombinationByMap(n int, numbers []int) int {
	intMap := make(map[int]bool)
	for _, v := range numbers {
		intMap[v] = true
	}

	return CalcSomeBy(n, intMap)
}

func CalcSomeBy(n int, mapping IntMap) int {
	fmt.Printf("n = %v mapping = %v\n", n, mapping)
	if n < 0 {
		return 0
	}
	total := 0
	intMap := make(IntMap)
	for k, v := range mapping {
		intMap[k] = v
	}
	for k, v := range intMap {
		if !v {
			continue
		}
		//if init {
		fmt.Printf("init %v \n", k)
		intMap[k] = false
		count := n / k
		fmt.Printf("count %v \n", count)
		i := 1
		for i <= count {
			fmt.Printf("process case with %v \n", i)
			diff := n - i*k
			if diff < 0 {
				continue
			}
			calcSomeBy := CalcSomeBy(diff, intMap)
			if calcSomeBy == 0 {
				calcSomeBy = 1
			}
			if diff == 0 {
				total += 1 * calcSomeBy
			} else {
				c2 := C2(diff + 1)
				fmt.Printf("diff = %v c2 = %v\n", diff, c2)
				total += c2 * calcSomeBy
			}
			fmt.Printf("total is %v\n", total)
			i++
		}
	}

	return total
}

/*
s {z1, z2, z3, ..., zk}

s/z1 = k1
s/z2 = k2
s/z3 = k3
...
s/zn = kn


1 -> C12 1
2 -> C22 2! 1
3 -> C32 3! 1

c2n


*/

func main() {
	//fmt.Println(C2(1))
	//fmt.Println(C2(2))
	//fmt.Println(C2(3))
	//fmt.Println(C2(4))
	//fmt.Println(C2(5))
	//fmt.Println(fibonacci(2))
	//fmt.Println(fibonacci(3))
	//fmt.Println(fibonacci(4))
	//fmt.Println(fibonacci(5))

	fmt.Printf("# %v\n ----\n", CombinationByMap(1, []int{1, 2}))
	fmt.Printf("# %v\n ----\n", CombinationByMap(2, []int{1, 2}))
	fmt.Printf("# %v\n ----\n", CombinationByMap(3, []int{1, 2}))
	//fmt.Printf("# %v\n ----\n", CombinationByMap(4, []int{1, 2, 3}))
	//fmt.Printf("# %v\n ----\n", CombinationByMap(5, []int{1, 2, 3}))
	//fmt.Printf("# %v\n ----\n", CombinationByMap(6, []int{1, 2, 3}))
}
