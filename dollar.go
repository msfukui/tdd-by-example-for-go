package money

// Dollar struct is the difinition of a US currency unit.
type Dollar struct {
	Money
}

// NewDollar is Dollar's constructor
func NewDollar(amount int) *Dollar {
	c := new(Dollar)
	c.amount = amount
	c.name = "Dollar"
	return c
}

// Times multiplies a Dollar amount by the specified value.
func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}
