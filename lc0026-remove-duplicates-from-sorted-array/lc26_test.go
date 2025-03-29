package lc0026_remove_duplicates_from_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantNums []int
		wantK    int
	}{
		{
			name:     "lc example 1",
			nums:     []int{1, 1, 2},
			wantNums: []int{1, 2},
			wantK:    2,
		},
		{
			name:     "lc example 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantNums: []int{0, 1, 2, 3, 4},
			wantK:    5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.nums) // Calls your implementation

			assert.Equal(t, tt.wantK, k)
			assert.Equal(t, tt.wantNums, tt.nums[:k])
		})
	}
}
