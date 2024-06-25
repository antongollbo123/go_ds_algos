package bubble_sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Unsorted", []int{3, 5, 1, 4, 2}, []int{1, 2, 3, 4, 5}},
		{"Duplicates", []int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{"Single element", []int{1}, []int{1}},
		{"Empty list", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := append([]int(nil), tt.input...) // create a copy of the input
			Sort(&list)
			if len(list) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(list))
			}
			for i := range list {
				if list[i] != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected, list)
					break
				}
			}
		})
	}
}
