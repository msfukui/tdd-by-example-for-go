package money

// Franc struct is the difinition of a US currency unit.
type Franc struct {
	Money
}

// NewFranc is Franc's constructor
func NewFranc(amount int) *Franc {
	c := new(Franc)
	c.amount = amount
	c.name = "Franc"
	return c
}

// Times multiplies a Franc amount by the specified value.
func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount * multiplier)
}
