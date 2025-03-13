package lc0027_remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantNums []int
		wantVal  int
	}{
		{
			name:     "lc test case 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantNums: []int{2, 2},
			wantVal:  2,
		},
		{
			name:     "lc test case 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantNums: []int{0, 1, 4, 0, 3},
			wantVal:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeElement(tt.nums, tt.val)

			assert.Equal(t, tt.wantVal, k)
			assert.Equal(t, tt.wantNums, tt.nums[:k])
		})
	}
}
