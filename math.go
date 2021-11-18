package money

// Sum returns the sum of all given values.
// The currencies must not differ.
//
//	Sum(MustFromString("12.34 ISO4217-EUR"), MustFromString("12.34 ISO4217-EUR")) // Returns 24.68 ISO4217-EUR.
//	Sum(MustFromString("12.34 ISO4217-EUR"), MustFromString("12.34"))             // Returns an error.
func Sum(first Value, values ...Value) (Value, error) {
	total := first.amount
	for _, value := range values {
		if first.currency != value.currency {
			return Value{}, &ErrorDifferentCurrencies{first.currency, value.currency}
		}
		total = total.Add(value.amount)
	}

	return Value{total, first.currency}, nil
}

// MustSum returns the sum of all given values.
// The currencies must not differ.
//
// Use this version if you have already made sure that the currencies are equal between all values.
func MustSum(first Value, values ...Value) Value {
	res, err := Sum(first, values...)
	if err != nil {
		panic(err)
	}
	return res
}
