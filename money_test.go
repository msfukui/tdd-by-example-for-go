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
	var result = NewDollar(5).Equals(NewDollar(5))

	if !result {
		t.Errorf("Dollar(%v)#Equals(%v) = %v, want %v", 5, 5, result, true)
	}

	result = NewDollar(5).Equals(NewDollar(6))

	if result {
		t.Errorf("Dollar(%v)#Equals(%v) = %v, want %v", 5, 6, result, false)
	}

	result = NewFranc(5).Equals(NewFranc(5))

	if !result {
		t.Errorf("Franc(%v)#Equals(%v) = %v, want %v", 5, 5, result, true)
	}

	result = NewFranc(5).Equals(NewFranc(6))

	if result {
		t.Errorf("Franc(%v)#Equals(%v) = %v, want %v", 5, 6, result, false)
	}

	result = NewFranc(5).Equals(NewDollar(5))

	if result {
		t.Errorf("Franc(%v)#Equals(Dollar(%v)) = %v, want %v", 5, 5, result, false)
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
