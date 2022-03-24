package main

import "fmt"

func swapTest() {
	n := []int{3, 4}
	swap(n, 0, 1)
	fmt.Println(n)
}

func swap(num []int, i, j int) {
	if i == j {
		return
	}
	num[i] = num[i] ^ num[j]
	num[j] = num[i] ^ num[j]
	num[i] = num[i] ^ num[j]
}
