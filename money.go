package money

// Money is a Dollar, Franc, or other base struct.
type Money struct {
	amounter int
}

// Moneyable is a Dollar, Franc, or other base interface.
type Moneyable interface {
	amount() int
}

func (d *Money) amount() int {
	return d.amounter
}

// Equals determines that the value specified in the argument is equal to
// the value of the receiver.
// If they are equal, return true.
func (d *Money) Equals(money Moneyable) bool {
	return d.amount() == money.amount()
}
