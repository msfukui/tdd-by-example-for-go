package money

// Franc struct is the difinition of a US currency unit.
type Franc struct {
	Money
}

// NewFranc is Franc's constructor
func NewFranc(amounter int) *Franc {
	c := new(Franc)
	c.amounter = amounter
	return c
}

// Times multiplies a Franc amount by the specified value.
func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount() * multiplier)
}

func (d *Franc) amount() int {
	return d.amounter
}
