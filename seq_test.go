/* seq_test.go

   Tests for graphic sequence functions. */

package graph

import (
	"testing"
)

func TestGraphic(t *testing.T) {
	tests := []struct {
		DegSeq
		bool
	}{
		{[]int{}, true},  // the empty graph
		{[]int{0}, true}, // an isolate
		{[]int{1}, false},
		{[]int{2}, false},
		{[]int{3}, false},
		{[]int{0, 0}, true},
		{[]int{1, 0}, false},
		{[]int{1, 1}, true}, // an edge
		{[]int{2, 1}, false},
		{[]int{2, 2}, false},
		{[]int{0, 0, 0}, true},
		{[]int{1, 0, 0}, false},
		{[]int{2, 0, 0}, false},
		{[]int{3, 0, 0}, false},
		{[]int{1, 1, 0}, true}, // edge + isolate
		{[]int{1, 1, 1}, false},
		{[]int{2, 1, 0}, false},
		{[]int{2, 1, 1}, true}, // P_3
		{[]int{2, 2, 1}, false},
		{[]int{2, 2, 2}, true}, // K_2
		{[]int{3, 1, 0}, false},
		{[]int{3, 1, 1}, false},
		{[]int{3, 2, 1}, false},
		{[]int{3, 2, 2}, false},
		{[]int{3, 3, 2}, false},
		{[]int{3, 3, 3}, false},
		{[]int{0, 0, 0, 0}, true},
		{[]int{1, 1, 0, 0}, true},
		{[]int{1, 1, 1, 1}, true},
		{[]int{2, 1, 1, 0}, true},
		{[]int{2, 2, 1, 1}, true}, // P_4
		{[]int{2, 2, 2, 0}, true},
		{[]int{2, 2, 2, 2}, true}, // C_4
		{[]int{3, 1, 1, 1}, true}, // claw
		{[]int{3, 2, 2, 1}, true},
		{[]int{3, 3, 3, 3}, true}, // K_4
		{[]int{1, 0, 0, 0}, false},
		{[]int{1, 1, 1, 0}, false},
		{[]int{2, 2, 1, 0}, false},
		{[]int{2, 2, 2, 1}, false},
		{[]int{3, 2, 2, 2}, false},
		{[]int{4, 2, 2, 2}, false},
		{[]int{4, 3, 3, 3}, false},
		{[]int{4, 4, 3, 3}, false},
		{[]int{4, 4, 4, 3}, false},
		{[]int{4, 4, 4, 4}, false},
		{[]int{3, 3, 3, 3, 3, 3}, true}, // K_3,3
	}

	for _, test := range tests {
		gr := test.DegSeq.Graphic()

		if gr != test.bool {
			t.Errorf("Graphic(%v) returned %t\n", test.DegSeq, gr)
		}
	}
}
