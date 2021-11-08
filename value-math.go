package d3money

// Equal returns if a monetary value is equal to another.
// If the currency is different between the values, the result will always be false.
func (v Value) Equal(comp Value) bool {
	if v.currency != comp.currency {
		return false
	}
	return v.value.Equal(comp.value)
}
