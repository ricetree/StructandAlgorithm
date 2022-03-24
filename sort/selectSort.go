package main

import "fmt"

func selectSortTest() {
	n := []int{1, 0, 3, 6, 7, 9, 8, 9, 22}
	selectSort(n)
	fmt.Println(n)
}

func selectSort(num []int) {
	ln := len(num)
	for i := 0; i < ln-1; i++ {
		min := i
		for j := i + 1; j < ln; j++ {
			if num[j] < num[i] {
				min = j
			}
		}
		swap(num, i, min)
	}
}
