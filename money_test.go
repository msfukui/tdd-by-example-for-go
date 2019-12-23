package money

import "testing"

func TestMultiplication(t *testing.T) {
	var five = NewDollar(5)

	var in, expected = 2, 10
	var product = five.Times(in)

	if x := product.amount; expected != x {
		t.Errorf("Dollar#Times(%v) = %v, want %v", in, x, expected)
	}

	in, expected = 3, 15
	product = five.Times(in)

	if x := product.amount; expected != x {
		t.Errorf("Dollar#Times(%v) = %v, want %v", in, x, expected)
	}
}

func TestEquality(t *testing.T) {
	var five = NewDollar(5)

	var in = 5
	var expected = true
	var result = five.Equals(NewDollar(in))
	if expected != result {
		t.Errorf("Dollar(%v)#Equals(%v) = %v, want %v", 5, in, result, expected)
	}

	in = 6
	expected = false
	result = five.Equals(NewDollar(in))
	if expected != result {
		t.Errorf("Dollar(%v)#Equals(%v) = %v, want %v", 5, in, result, expected)
	}
}
