package main

import (
	"fmt"
)

func main() {
	start := []int{0,14,6,20,1,4}
	prev := 0

	m := map[int][]int{}
	for i, v := range start {
		m[v] = []int{i}
		prev = v
	}
	test(start)

	println("============")
	for i := len(start); i < 30000000; i++ {
		last, found := m[prev]
		//fmt.Println("Looking for", prev, found, last)
		if found && len(last) > 1 {
			prev = i - last[len(last) - 2] - 1 
		} else {
			prev = 0
		}
		m[prev] = append(m[prev], i)
		//fmt.Println(i, "->", prev)
	}
	println(prev)
}


func test(start []int) {
	nums := make([]int, 20)
	for i, _ := range nums {
		if i < len(start) {
			nums[i] = start[i]
		} else {
			last := indexOf(nums, nums[i-1], i-1)
			if last == -1 {
				nums[i] = 0
			} else {
				nums[i] = (i-1) - last
			}
		}
		println(i, nums[i])
	}
	fmt.Println(nums[len(nums) - 1])
}

func indexOf(arr []int, val int, max int) int {
	l := max - 1
	for l >= 0 {
		if arr[l] == val {
			return l
		}
		l--
	}
	return -1
}

