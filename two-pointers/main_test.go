package main

import (
	"reflect"
	"testing"
)

func TestFindSumInSortedArray(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 6},
			target: 6,
			output: []int{1, 3},
		},
		{
			input:  []int{3, 4, 5, 7, 9},
			target: 14,
			output: []int{2, 4},
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6},
			target: 13,
			output: []int{},
		},
	}

	for _, v := range tests {
		assertEquals(t, findSumInSortedArray(v.input, v.target), v.output)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{2, 3, 3, 3, 6, 9, 9},
			output: 4,
		},
		{
			input:  []int{2, 2, 2, 11},
			output: 2,
		},
	}
	for _, v := range tests {
		assertEquals(t, removeDuplicates(v.input), v.output)
	}
}

func TestRemoveKey(t *testing.T) {
	tests := []struct {
		input  []int
		key    int
		output int
	}{
		{
			input:  []int{3, 2, 3, 6, 3, 10, 9, 3},
			key:    3,
			output: 4,
		},
		{
			input:  []int{2, 11, 2, 2, 1},
			key:    2,
			output: 2,
		},
	}
	for _, v := range tests {
		assertEquals(t, removeKey(v.input, v.key), v.output)
	}
}

func TestSquareSortedArray(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{-4, -1, 0, 3, 10},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			input:  []int{-7, -3, 2, 3, 11},
			output: []int{4, 9, 9, 49, 121},
		},
		{
			input:  []int{-5, -4, -3, -2, -1},
			output: []int{1, 4, 9, 16, 25},
		},
	}

	for _, v := range tests {
		assertEquals(t, squareSortedArray(v.input), v.output)
	}
}

func TestSearchTriplets(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{-3, 0, 1, 2, -1, 1, -2},
			output: [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}},
		},
		{
			input:  []int{-5, 2, -1, -2, 3},
			output: [][]int{{-5, 2, 3}, {-2, -1, 3}},
		},
	}

	for _, v := range tests {
		assertEquals(t, searchTriplets(v.input), v.output)
	}
}

func TestTripletSumClosestToTarget(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output int
	}{
		{
			input:  []int{-2, 0, 1, 2},
			target: 2,
			output: 1,
		},
		{
			input:  []int{-3, -1, 1, 2},
			target: 1,
			output: 0,
		},
		{
			input:  []int{1, 0, 1, 1},
			target: 100,
			output: 3,
		},
	}

	for _, v := range tests {
		assertEquals(t, tripletSumClosestToTarget(v.input, v.target), v.output)
	}
}

func TestTripletSumSmallerThanTarget(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output int
	}{
		{
			input:  []int{-1, 0, 2, 3},
			target: 3,
			output: 2,
		},
		{
			input:  []int{-1, 4, 2, 1, 3},
			target: 5,
			output: 4,
		},
		{
			input:  []int{},
			target: 0,
			output: 0,
		},
		{
			input:  []int{0},
			target: 0,
			output: 0,
		},
	}

	for _, v := range tests {
		assertEquals(t, tripletSumSmallerThanTarget(v.input, v.target), v.output)
	}
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output [][]int
	}{
		{
			input:  []int{1, 0, -1, 0, -2, 2},
			target: 0,
			output: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			input:  []int{},
			target: 0,
			output: [][]int{},
		},
	}

	for _, v := range tests {
		assertEquals(t, fourSum(v.input, v.target), v.output)
	}
}

func TestCompareString(t *testing.T) {
	tests := []struct {
		str1   string
		str2   string
		output bool
	}{
		{
			str1:   "xy#z",
			str2:   "xzz#",
			output: true,
		},
		{
			str1:   "xy#z",
			str2:   "xyz#",
			output: false,
		},
		{
			str1:   "xp#",
			str2:   "xyz##",
			output: true,
		},
		{
			str1:   "xywrrmp",
			str2:   "xywrrmu#p",
			output: true,
		},
	}

	for _, v := range tests {
		assertEquals(t, compareStrings(v.str1, v.str2), v.output)
	}
}

func assertEquals(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
