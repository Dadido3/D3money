// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Equal returns if a monetary value is equal to another.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) Equal(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.Equal(comp.amount), nil
}

// GreaterThan returns if the monetary value is greater than another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) GreaterThan(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.GreaterThan(comp.amount), nil
}

// GreaterThanOrEqual returns if the monetary value is greater than or equal to another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) GreaterThanOrEqual(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.GreaterThanOrEqual(comp.amount), nil
}

// LessThan returns if the monetary value is less than another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) LessThan(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.LessThan(comp.amount), nil
}

// LessThanOrEqual returns if the monetary value is less than or equal to another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) LessThanOrEqual(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.LessThanOrEqual(comp.amount), nil
}

// Decimal returns the value as a shopspring/decimal number.
func (v Value) Decimal() decimal.Decimal {
	return v.amount
}

// Float64 returns the nearest float64 for the value v, and a bool indicating if the float represents the value exactly.
func (v Value) Float64() (f float64, exact bool) {
	return v.amount.Float64()
}

// InexactFloat64 returns the nearest float64 for the value v.
func (v Value) InexactFloat64() float64 {
	f, _ := v.amount.Float64()
	return f
}

// Add returns v + v2 as a new value.
// It will not mutate either v or v2.
//
// In case the two values don't use the same currency, this will return an error.
func (v Value) Add(v2 Value) (Value, error) {
	if v.currency != v2.currency {
		return Value{}, &ErrorDifferentCurrencies{v.currency, v2.currency}
	}

	return Value{amount: v.amount.Add(v2.amount), currency: v.currency}, nil
}

// MustAdd returns v + v2 as a new value.
// It will not mutate either v or v2.
//
// In case the two values don't use the same currency, this will panic.
// Use this version if you have already made sure that the currency is equal between both values.
func (v Value) MustAdd(v2 Value) Value {
	res, err := v.Add(v2)
	if err != nil {
		panic(err)
	}
	return res
}

// Sub returns v - v2 as a new value.
// It will not mutate either v or v2.
//
// In case the two values don't use the same currency, this will return an error.
func (v Value) Sub(v2 Value) (Value, error) {
	if v.currency != v2.currency {
		return Value{}, &ErrorDifferentCurrencies{v.currency, v2.currency}
	}

	return Value{amount: v.amount.Sub(v2.amount), currency: v.currency}, nil
}

// MustSub returns v - v2 as a new value.
// It will not mutate either v or v2.
//
// In case the two values don't use the same currency, this will panic.
// Use this version if you have already made sure that the currency is equal between both values.
func (v Value) MustSub(v2 Value) Value {
	res, err := v.Sub(v2)
	if err != nil {
		panic(err)
	}
	return res
}

// Mul returns v * v2 as a new value.
// It will not mutate either v or v2.
//
// In case both values have a currency, this will return an error.
func (v Value) Mul(v2 Value) (Value, error) {
	var currency Currency

	if v.currency != nil && v2.currency != nil {
		return Value{}, fmt.Errorf("can't multiply two values with currencies %s and %s", helperCurrencyUniqueCode(v.currency), helperCurrencyUniqueCode(v2.currency))
	} else if v.currency != nil {
		currency = v.currency
	} else if v2.currency != nil {
		currency = v2.currency
	}

	return Value{amount: v.amount.Mul(v2.amount), currency: currency}, nil
}

// MustMul returns v * v2 as a new value.
// It will not mutate either v or v2.
//
// In case both values have a currency, this will panic.
func (v Value) MustMul(v2 Value) Value {
	res, err := v.Mul(v2)
	if err != nil {
		panic(err)
	}
	return res
}

// Abs returns the absolute value.
//
//	Abs(FromString("-123.456 ISO4217-EUR")) // Returns "123.456 ISO4217-EUR"
//	Abs(FromString("123.456 ISO4217-EUR"))  // Returns "123.456 ISO4217-EUR"
func (v Value) Abs() Value {
	return Value{amount: v.amount.Abs(), currency: v.currency}
}

// Neg returns the negative value.
//
//	Neg(FromString("-123.456 ISO4217-EUR")) // Returns "123.456 ISO4217-EUR"
//	Neg(FromString("123.456 ISO4217-EUR"))  // Returns "-123.456 ISO4217-EUR"
func (v Value) Neg() Value {
	return Value{amount: v.amount.Neg(), currency: v.currency}
}

// Sign returns:
//
//	-1 if v <  0
//	 0 if v == 0
//	+1 if v >  0
//
// The currency is ignored.
func (v Value) Sign() int {
	return v.amount.Sign()
}

// IsPositive returns true when the value is greater than zero, false otherwise.
// The currency is ignored.
func (v Value) IsPositive() bool {
	return v.Sign() == 1
}

// IsNegative returns true when the value is smaller than zero, false otherwise.
// The currency is ignored.
func (v Value) IsNegative() bool {
	return v.Sign() == -1
}

// IsZero returns true when the value is exactly zero.
// The currency is ignored.
func (v Value) IsZero() bool {
	return v.Sign() == 0
}
