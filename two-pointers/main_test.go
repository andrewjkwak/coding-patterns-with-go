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
		assertArrayEquals(t, findSumInSortedArray(v.input, v.target), v.output)
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

func assertArrayEquals(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
