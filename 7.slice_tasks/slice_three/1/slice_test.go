package main

import (
	"reflect"
	"testing"
)

func TestRemoveUnordered(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{"remove middle", []int{1, 2, 3, 4}, 1, []int{1, 4, 3}},
		{"remove first", []int{1, 2, 3}, 0, []int{3, 2}},
		{"remove last", []int{1, 2, 3}, 2, []int{1, 2}},
		{"invalid index", []int{1, 2, 3}, 10, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveUnordered(append([]int(nil), tt.input...), tt.index)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRemoveOrdered(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		index    int
		expected []string
	}{
		{"remove middle", []string{"a", "b", "c"}, 1, []string{"a", "c"}},
		{"remove first", []string{"x", "y", "z"}, 0, []string{"y", "z"}},
		{"invalid index", []string{"a"}, 5, []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveOrdered(append([]string(nil), tt.input...), tt.index)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRemoveAllByValue(t *testing.T) {
	input := []int{1, 2, 3, 2, 4}
	expected := []int{1, 3, 4}
	result := RemoveAllByValue(input, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []string{"go", "go", "rust"}
	expected := []string{"go", "rust"}
	result := RemoveDuplicates(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestRemoveIf(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 3, 5}
	result := RemoveIf(input, func(n int) bool { return n%2 == 0 })
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestShrinkCapacity(t *testing.T) {
	// Создаём слайс с большим cap
	s := make([]int, 5, 20)
	result := ShrinkCapacity(s)
	if cap(result) != len(result) {
		t.Errorf("expected cap=%d, got %d", len(result), cap(result))
	}
}
