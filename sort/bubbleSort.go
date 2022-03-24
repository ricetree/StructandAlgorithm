package main

import "fmt"

func bubbleSortTest() {
	n := []int{1, 0, 3, 6, 7, 9, 8, 9, 22}
	BubbleSort(n)
	fmt.Println(n)
}

func BubbleSort(num []int) {
	ln := len(num)
	for i := ln - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if num[j] > num[j+1] {
				swap(num, j, j+1)
			}
		}
	}
}
