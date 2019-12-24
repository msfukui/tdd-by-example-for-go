package money

import "testing"

func TestMultiplication(t *testing.T) {
	var five = NewDollar(5)

	var result = five.Times(2).Equals(NewDollar(10))

	if !result {
		t.Errorf("Dollar#Times(%v) == %v is %v, want %v", 2, NewDollar(10), true, result)
	}

	result = five.Times(3).Equals(NewDollar(15))

	if !result {
		t.Errorf("Dollar#Times(%v) == %v is %v, want %v", 3, NewDollar(15), true, result)
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

func TestFrancMultiplication(t *testing.T) {
	var five = NewFranc(5)

	var result = five.Times(2).Equals(NewFranc(10))

	if !result {
		t.Errorf("Franc#Times(%v) == %v is %v, want %v", 2, NewFranc(10), true, result)
	}

	result = five.Times(3).Equals(NewFranc(15))

	if !result {
		t.Errorf("Franc#Times(%v) == %v is %v, want %v", 3, NewFranc(15), true, result)
	}
}
