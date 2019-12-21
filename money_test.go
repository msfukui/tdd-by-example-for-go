package money

import "testing"

func TestMultiplication(t *testing.T) {
	var five = Dollar{5}
	five.times(2)

	if x := five.amount; 10 != x {
		t.Errorf("Dollar#time(%v) = %v, want %v", 2, x, 10)
	}
}
