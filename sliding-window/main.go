package main

import (
	"fmt"
	"strings"
) 

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinString(a, b string) string {
	if len(a) < len(b) {
		return a
	}
	return b
}

// Given an array, find the average of all contiguous subarrays of size k.
func findAverageOfSubarrayK(nums []int, k int) []float32 {
	if len(nums) < k {
		return []float32{}
	}

	out := []float32{}
	sum := 0
	for _, v := range nums[0:5] {
		sum += v
	}
	out = append(out, float32(sum)/float32(k))

	for i := 1; i <= len(nums)-k; i++ {
		sum = sum - nums[i-1] + nums[i+k-1]
		out = append(out, float32(sum)/float32(k))
	}

	return out
}

// Given an array of postive numbers and positive number k, find the max sum of any contiguous subarray of size k.
func findMaxSumSubarray(nums []int, k int) int {
	sum := 0
	// get the initial max calue
	for _, v := range nums[0:k] {
		sum += v
	}
	max := sum
	for i := 1; i <= len(nums)-k; i++ {
		sum = sum - nums[i-1] + nums[i+k-1]
		max = getMax(max, sum)
	}
	return max
}

// Given an array of postive numbers and a positive number S, find the length of the smallest contiguous
// subarray whose sum is greater than or equal to S.  Return 0, if no such subarray exists.
func findMinSubarray(nums []int, S int) int {
	res := len(nums)
	sum := 0
	for start, end := 0, 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= S {
			res = getMin(res, end-start+1)
			sum -= nums[start]
			start++
		}
	}
	if res == len(nums) {
		return 0
	}
	return res
}

// Given a string, find the length of the longest substring in it with no more than K distinct characters.
func longestSubstringWithKDistinct(str []byte, K int) int {
	start, max := 0, 0
	count := make(map[byte]int)

	for end := 0; end < len(str); end++ {
		count[str[end]]++

		for len(count) > K {
			count[str[start]]--
			if count[str[start]] == 0 {
				delete(count, str[start])
			}
			start++
		}
		max = getMax(max, end-start+1)
	}
	return max
}

// Given an array of characters where each character represents a fruit tree, you are given two baskets
// and your goal is to put max number of fruits in each basket.  The only restriction is that each basket
// can have only one type of fruit.
func fruitsIntoBaskets(str []byte) int {
	start, max := 0, 0
	count := make(map[byte]int)

	for end := 0; end < len(str); end++ {
		count[str[end]]++

		for len(count) > 2 {
			count[str[start]]--
			if count[str[start]] == 0 {
				delete(count, str[start])
				start++
			}
		}
		max = getMax(max, end-start+1)
	}
	return max
}

// Given a string, find the length of the longest substring which has no repeating characters.
func nonRepeatSubstring(str []byte) int {
	start, max := 0, 0
	set := make(map[byte]bool)

	// add each element into the "set"
	for _, char := range str {
		set[char] = false
	}

	for end := 0; end < len(str); end++ {
		exists := set[str[end]]
		// if it exists, that means we hit a duplicate
		if exists {
			for set[str[end]] {
				set[str[start]] = false
				start++
			}
		}
		set[str[end]] = true
		max = getMax(max, end-start+1)
	}
	return max
}

// Given a string with lowercase letters only, if you are allowed to replace no more than 'k'
// letters with any letter, find the length of the longest substring having the same letters after replacement
func findLongestSubstringReplaceK(str []byte, k int) int {
	var start, max, maxRepeat int
	count := make(map[byte]int)

	for end := 0; end < len(str); end++ {
		count[str[end]]++
		maxRepeat = getMax(maxRepeat, count[str[end]])

		if end-start+1-maxRepeat > k {
			count[str[start]]--
			start++
		}
		max = getMax(max, end-start+1)
	}
	return max
}

// Given a string and a pattern, find out if the string contains any permutation of the pattern.
func permutationPatternInString(str, pattern string) bool {
	patternCount := make(map[byte]int)
	// Add all the chars from pattern into patternCount
	for i := 0; i < len(pattern); i++ {
		patternCount[pattern[i]]++
	}
	matched, start := 0, 0
	// Iterate thru str, and every single time patternCount of char hits 0,
	// we know we fulfilled one of the chars and we can increment matched
	// Once matched hits len(patternCount), we know we have a match and can return true
	// We have to make sure our "window" is only to len(pattern)... If we go above it,
	// increment start by 1 to slide the window down
	for end := 0; end < len(str); end++ {
		_, ok := patternCount[str[end]]
		// We only care if the char exists in patternCount
		if ok {
			patternCount[str[end]]--
			if patternCount[str[end]] == 0 {
				matched++
			}
		}
		if matched == len(patternCount) {
			return true
		}
		// If our window size is too big, slide it
		if end-start+1 >= len(pattern) {
			_, ok := patternCount[str[start]]
			// if the start-index char exists in pattern count, we want to increment it since we're basically
			// "putting it back" and we want to decrement matched if the count was 0 since it's now unmatched
			if ok {
				if patternCount[str[start]] == 0 {
					matched--
				}
				patternCount[str[start]]++
			}
			start++
		}
	}
	return false
}

// Given a string and a pattern, find all anagrams of the pattern in the given string.
// ex: str: "abbcabc", pattern: "abc" -> anagram from index 2, index 3, and index 4 [2, 3, 4]
func findStringAnagrams(str, pattern string) []int {
	// This is similar to last problem, except we get all instances of it and return array of indices
	patternCount := make(map[byte]int)
	res := []int{}
	for i := 0; i < len(pattern); i++ {
		patternCount[pattern[i]]++
	}
	matched, start := 0, 0
	for end := 0; end < len(str); end++ {
		_, ok := patternCount[str[end]]
		if ok {
			patternCount[str[end]]--
			if patternCount[str[end]] == 0 {
				matched++
			}
		}
		// If we found a match, add the start index to the slice and decrement matched again
		if matched == len(patternCount) {
			res = append(res, start)
		}
		// Slide the window by 1 if we reached our window size
		if end-start+1 >= len(pattern) {
			_, ok := patternCount[str[start]]
			if ok {
				if patternCount[str[start]] == 0 {
					matched--
				}
				patternCount[str[start]]++
			}
			start++
		}
	}
	return res
}

// Given a string and a pattern, find the smallest substring in the given string which has 
// all the characters of the given pattern.
// ex: str: "aabdec", pattern: "abc" -> "abdec" since c is at the end
func minimumWindowSubstring(str, pattern string) string {
	patternCount := make(map[byte]int)
	for i := range pattern {
		patternCount[pattern[i]]++
	}

	matched, start := 0, 0
	res := string(make([]byte, len(str)))
	for end := 0; end < len(str); end++ {
		if _, ok := patternCount[str[end]]; ok {
			patternCount[str[end]]--
			if patternCount[str[end]] == 0 {
				matched++
			}
		}
		for matched >= len(patternCount) {
			if len(res) >= len(str[start:end+1]) {
				res = str[start:end+1]
			}
			if _, ok := patternCount[str[start]]; ok {
				patternCount[str[start]]++
				if patternCount[str[start]] > 0 {
					matched--
				}
			}
			start++
		}
	}
	if res == string(make([]byte, len(str))) {
		return ""
	}
	return res
}

// Give a string and a list of words, find all the starting indices of substrings in the given string
// that are a concatenation of all the given words exactly once without any overlapping of words.
// It is a given that all words are of the same length.
func findWordConcatenation(str string, words []string) []int {
	
}

func main() {
}
