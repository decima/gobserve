package gobserve

import "testing"

func TestSortIntMap(t *testing.T) {
	input := map[int]int{0: 0, 10: 10, 7: 7, 8: 8}
	expectedOutput := []int{10, 8, 7, 0}

	output := sortIntMap(input)

	if len(expectedOutput) != len(output) {
		t.Errorf("Expected String(%v) is not same as"+
			" actual string (%v)", expectedOutput, output)
	}
	for i := 0; i < len(output); i++ {

		if expectedOutput[i] != output[i] {
			t.Errorf("Expected slice (%v) is not same as"+
				" actual slice (%v)", expectedOutput, output)
		}
	}

}
