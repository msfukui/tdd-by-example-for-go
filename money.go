package money

// Dollar struct is the difinition of a US currency unit.
type Dollar struct {
	amount int
}

// Times multiplies a Dollar amount by the specified value.
func (d *Dollar) Times(multiplier int) {
	d.amount *= multiplier
}
