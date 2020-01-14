package money

// Money is a Dollar, Franc, or other base struct.
type Money struct {
	amount int
	name   string
}

// Moneyable is a Dollar, Franc, or other base interface.
type Moneyable interface {
	getAmount() int
	getName() string
}

func (d *Money) getAmount() int {
	return d.amount
}

func (d *Money) getName() string {
	return d.name
}

// Equals determines that the value specified in the argument is equal to
// the value of the receiver.
// If they are equal, return true.
func (d *Money) Equals(object interface{}) bool {
	var money = object.(Moneyable)
	return d.amount == money.getAmount() && d.name == money.getName()
}
