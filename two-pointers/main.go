package main

import "fmt"

// Helper functions
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Given an array of sorted numbers and a target sum, find a pair in the array
// whose sum is equal to the given target.
func findSumInSortedArray(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			return []int{start, end}
		} else if sum > target {
			end--
		} else {
			start++
		}
	}
	return []int{}
}

// Given an array of sorted numbers, remove all duplicates from it.
// You should not use any extra space; after removing the duplicates in-place
// return the length of the subarray that has no duplicate in it.
// ex: [2, 3, 3, 3, 6, 9, 9] -> 4 b/c [2, 3, 6, 9]
func removeDuplicates(nums []int) int {
	next := 1
	for i := 1; i < len(nums); i++ {
		if nums[next-1] != nums[i] {
			nums[next] = nums[i]
			next++
		}
	}
	return next
}

// Given an unsorted array of numbers and a target key, remove all instance of key
// in-place and return the new length of the array
func removeKey(nums []int, key int) int {
	next := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != key {
			nums[next] = nums[i]
			next++
		}
	}
	return next
}

// Given a sorted array, create a new array containing squares of all the
// number of the input array in the sorted array
// func squareSortedArray(nums []int) []int {
// 	var p, n int
// 	res := []int{}
// 	for p < len(nums) {
// 		if nums[p] < 0 {
// 			p++
// 		}
// 	}
// 	return []int{}
// }

func main() {
	x := make([]int, _, 10)
	x = append(x, 1)
	fmt.Printf("%v", x)
}
