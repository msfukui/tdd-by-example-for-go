package money

// Franc struct is the difinition of a US currency unit.
type Franc struct {
	amount int
}

// NewFranc is Franc's constructor
func NewFranc(amount int) *Franc {
	c := new(Franc)
	c.amount = amount
	return c
}

// Times multiplies a Franc amount by the specified value.
func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount * multiplier)
}

// Equals determines that the value specified in the argument is equal to
// the value of the receiver.
// If they are equal, return true.
func (d *Franc) Equals(object *Franc) bool {
	return d.amount == object.amount
}
