package money

// Dollar struct is the difinition of a US currency unit.
type Dollar struct {
	amount int
}

// NewDollar is Dollar's constructor
func NewDollar(amount int) *Dollar {
	c := new(Dollar)
	c.amount = amount
	return c
}

// Times multiplies a Dollar amount by the specified value.
func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

// Equals determines that the value specified in the argument is equal to
// the value of the receiver.
// If they are equal, return true.
func (d *Dollar) Equals(object *Dollar) bool {
	return d.amount == object.amount
}
