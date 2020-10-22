package main

import (
	"math"
	"sort"
)

// Helper functions
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
func squareSortedArray(nums []int) []int {
	var p, n int
	res := []int{}
	// We need to find the first non-negative integer but also make sure we don't exit the slice
	for nums[p] < 0 && p < len(nums)-1 {
		p++
	}
	// At this point, we have the index for the first non-negative value
	// The one right before it is the pointer for the negative values
	// We compare n and p to see which value is less (in absolute value)
	// We also need to make sure that n >= 0 && p < len(nums)
	n = p - 1
	for n >= 0 && p < len(nums) {
		if Abs(nums[n]) < nums[p] {
			res = append(res, nums[n]*nums[n])
			n--
		} else {
			res = append(res, nums[p]*nums[p])
			p++
		}
	}
	for n >= 0 {
		res = append(res, nums[n]*nums[n])
		n--
	}
	for p < len(nums) {
		res = append(res, nums[p]*nums[p])
		p++
	}
	return res
}

// Given an array of unsorted numbers, find all unique triplets in it that add up to zero.
func searchTriplets(nums []int) [][]int {
	triplets := [][]int{}
	// Sort the array to be able to use two-pointer
	sort.Ints(nums)
	// Make sure we go len(nums)-2 before so that we have enough values after where a triplet could exist
	for i := 0; i < len(nums)-2; i++ {
		// We skip the duplicates (for 0, we don't have to since there's no number before it)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Get the target (the negative value of the value we're currently on)
		// left is the index after the number we're currently on, as we're searching after it
		// right is the end of the array
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			// if the sum is equal to the target, we found a triplet
			if sum == target {
				triplets = append(triplets, []int{nums[i], nums[left], nums[right]})
				// we still need to keep searching, as there might be other triplets available for that value
				// we make sure to prevent adding any duplicates using the while-loops
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return triplets
}

// Given an array of unsorted numbers and a target number, find a triplet in the array whose sum is
// as close to the target number as possible, return the sum of the triplet.
func tripletSumClosestToTarget(nums []int, target int) int {
	res := math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			threeSum := nums[i] + nums[left] + nums[right]
			// Check to see if we found a match, else get the threeSum with the lowest diff to target
			if target == threeSum {
				return threeSum
			} else if Abs(target-threeSum) < Abs(target-res) {
				res = threeSum
			}

			// We know if our threeSum is greater than target, moving our right up will just increase
			// that difference... taking us further away from the closest sum
			// Therefore, we decrement right by 1, else increment left by 1
			if threeSum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

// Given an array nums of unsorted numbers and a target sum, count all triplets in it such that
// the sum of the triplet is less than target.  Write a function to return the count of the triplets.
func tripletSumSmallerThanTarget(nums []int, target int) int {
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			threeSum := nums[i] + nums[left] + nums[right]
			// We know if threeSums is less than target, all numbers between start and end is also less
			// than target...
			if threeSum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}
	return count
}

// func subarrayProductLessThanK(nums []int, target int) [][]int {

// }

// Given an array containing 0s, 1s and 2s, sort the array in-place.
func dutchFlagSort(nums []int) {
	start, end := 0, len(nums)-1
	for i := 0; i <= end; {
		if nums[i] == 0 {
			nums[i], nums[start] = nums[start], nums[i]
			i++
			start++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		}
	}
}

// Given an array of numbers, find all instances of 4-sums that equals a target
func fourSum(nums []int, target int) [][]int {
	// This is basically threeSum with an extra step... I'll be creating a threeSum function
	// that's a little modified to help me solve this problem
	quad := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum(&quad, nums[i], nums[i+1:], target-nums[i])
	}
	return quad
}

// Helper function for fourSum
func threeSum(quad *[][]int, first int, nums []int, target int) {
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			threeSum := nums[i] + nums[left] + nums[right]
			if threeSum == target {
				*quad = append(*quad, []int{first, nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[left] == nums[right+1] {
					right--
				}
			} else if threeSum > target {
				right--
			} else {
				left++
			}
		}
	}
}

// Given two strings containing backspaces (identified by '#'), check if two strings are equal.
func compareStrings(str1, str2 string) bool {
	idx1, idx2 := len(str1)-1, len(str2)-1
	for idx1 >= 0 || idx2 >= 0 {
		i1 := nextValidChar(str1, idx1)
		i2 := nextValidChar(str2, idx2)

		if i1 < 0 && i2 < 0 {
			return true
		}

		if i1 < 0 || i2 < 0 {
			return false
		}

		if str1[i1] != str2[i2] {
			return false
		}

		idx1 = i1 - 1
		idx2 = i2 - 1
	}
	return true
}

// Helper function for compareStrings
func nextValidChar(str string, i int) int {
	backspaceCount := 0
	for i >= 0 {
		if str[i] == '#' {
			backspaceCount++
		} else if backspaceCount > 0 {
			backspaceCount--
		} else {
			break
		}
		i--
	}
	return i
}

func main() {
}
