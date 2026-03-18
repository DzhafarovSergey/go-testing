package even

import "testing"

func TestSuccessEven(t *testing.T) {
	num_even := IsEven(2)

	if num_even != true {
		t.Errorf("Number 2 should be even, got %v", num_even)
	}
}

func TestNotEven(t *testing.T) {
	num_even := IsEven(3)

	if num_even != false {
		t.Errorf("Number 3 should not be even, got %v", num_even)
	}
}
