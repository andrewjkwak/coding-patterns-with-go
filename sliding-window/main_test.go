package main

import (
	"reflect"
	"testing"
)

func TestAverageOfSubarrayK(t *testing.T) {
	input := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}

	got := findAverageOfSubarrayK(input, 5)
	want := []float32{2.2, 2.8, 2.4, 3.6, 2.8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxSumOfSubarrayK(t *testing.T) {
	input := []int{2, 1, 5, 1, 3, 2}
	got := findMaxSumSubarray(input, 3)
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFindMinSubarray(t *testing.T) {
	tests := []struct {
		input  []int
		S      int
		output int
	}{
		{
			input:  []int{2, 1, 5, 2, 3, 2},
			S:      7,
			output: 2,
		},
		{
			input:  []int{2, 1, 5, 2, 8},
			S:      7,
			output: 1,
		},
		{
			input:  []int{3, 4, 1, 1, 6},
			S:      8,
			output: 3,
		},
	}

	for _, v := range tests {
		assertEquals(t, findMinSubarray(v.input, v.S), v.output)
	}
}

func TestLongestSubstringWithKDistinct(t *testing.T) {
	tests := []struct {
		input  []byte
		K      int
		output int
	}{
		{
			input:  []byte{'a', 'r', 'a', 'a', 'c', 'i'},
			K:      2,
			output: 4,
		},
		{
			input:  []byte{'a', 'r', 'a', 'a', 'c', 'i'},
			K:      1,
			output: 2,
		},
		{
			input:  []byte{'c', 'b', 'b', 'e', 'b', 'i'},
			K:      3,
			output: 5,
		},
	}

	for _, v := range tests {
		assertEquals(t, longestSubstringWithKDistinct(v.input, v.K), v.output)
	}
}

func TestFruitsIntoBaskets(t *testing.T) {
	tests := []struct {
		input  []byte
		output int
	}{
		{
			input:  []byte{'A', 'B', 'C', 'A', 'C'},
			output: 3,
		},
		{
			input:  []byte{'A', 'B', 'C', 'B', 'B', 'C'},
			output: 5,
		},
	}

	for _, v := range tests {
		assertEquals(t, fruitsIntoBaskets(v.input), v.output)
	}
}

func TestNonRepeatSubstring(t *testing.T) {
	tests := []struct {
		input  []byte
		output int
	}{
		{
			input:  []byte{'a', 'b', 'c', 'a', 'b', 'c', 'b', 'b'},
			output: 3,
		},
		{
			input:  []byte{'b', 'b', 'b', 'b', 'b'},
			output: 1,
		},
		{
			input:  []byte{'p', 'w', 'w', 'k', 'e', 'w'},
			output: 3,
		},
		{
			input:  []byte{},
			output: 0,
		},
	}

	for _, v := range tests {
		assertEquals(t, nonRepeatSubstring(v.input), v.output)
	}
}

func TestFindLongestSubstringReplaceK(t *testing.T) {
	tests := []struct {
		input  []byte
		k      int
		output int
	}{
		{
			input:  []byte{'a', 'a', 'b', 'c', 'c', 'b', 'b'},
			k:      2,
			output: 5,
		},
		{
			input:  []byte{'a', 'b', 'b', 'c', 'b'},
			k:      1,
			output: 4,
		},
		{
			input:  []byte{'a', 'b', 'c', 'c', 'd', 'e'},
			k:      1,
			output: 3,
		},
	}

	for _, v := range tests {
		assertEquals(t, findLongestSubstringReplaceK(v.input, v.k), v.output)
	}
}

func assertEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
