package lc0088_merge_sorted_array

import "testing"
import "github.com/stretchr/testify/assert"

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "lc example 1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "lc example 2",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "lc example 3",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
		{
			name:  "lc example 4",
			nums1: []int{4, 0, 0, 0, 0, 0},
			m:     1,
			nums2: []int{1, 2, 3, 5, 6},
			n:     5,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.want, tt.nums1)
		})
	}
}
