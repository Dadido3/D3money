// Copyright (c) 2021-2023 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

// Equal returns if a monetary value is equal to another.
// If the currency differs between the two values, the result is always false.
func (v Value) Equal(comp Value) bool {
	if v.currency != comp.currency {
		return false
	}
	return v.amount.Equal(comp.amount)
}

// EqualDetailed returns if a monetary value is equal to another.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) EqualDetailed(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.Equal(comp.amount), nil
}

// GreaterThan returns if the monetary value is greater than another value.
// If the currency differs between the two values, the result is always false.
func (v Value) GreaterThan(comp Value) bool {
	if v.currency != comp.currency {
		return false
	}
	return v.amount.GreaterThan(comp.amount)
}

// GreaterThanDetailed returns if the monetary value is greater than another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) GreaterThanDetailed(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.GreaterThan(comp.amount), nil
}

// GreaterThanOrEqual returns if the monetary value is greater than or equal to another value.
// If the currency differs between the two values, the result is always false.
func (v Value) GreaterThanOrEqual(comp Value) bool {
	if v.currency != comp.currency {
		return false
	}
	return v.amount.GreaterThanOrEqual(comp.amount)
}

// GreaterThanOrEqualDetailed returns if the monetary value is greater than or equal to another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) GreaterThanOrEqualDetailed(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.GreaterThanOrEqual(comp.amount), nil
}

// LessThan returns if the monetary value is less than another value.
// If the currency differs between the two values, the result is always false.
func (v Value) LessThan(comp Value) bool {
	if v.currency != comp.currency {
		return false
	}
	return v.amount.LessThan(comp.amount)
}

// LessThanDetailed returns if the monetary value is less than another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) LessThanDetailed(comp Value) (bool, error) {
	if v.currency != comp.currency {
		return false, &ErrorDifferentCurrencies{v.currency, comp.currency}
	}
	return v.amount.LessThan(comp.amount), nil
}

// LessThanOrEqual returns if the monetary value is less than or equal to another value.
// If the currency differs between the two values, the result is always false.
func (v Value) LessThanOrEqual(comp Value) bool {
	if v.currency != comp.currency {
		return false
	}
	return v.amount.LessThanOrEqual(comp.amount)
}

// LessThanOrEqualDetailed returns if the monetary value is less than or equal to another value.
// If the currency differs between the two values, the result is always false and an error is returned.
func (v Value) LessThanOrEqualDetailed(comp Value) (bool, error) {
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
//	MustFromString("-123.456 ISO4217-EUR").Abs() // Returns "123.456 ISO4217-EUR"
//	MustFromString("123.456 ISO4217-EUR").Abs()  // Returns "123.456 ISO4217-EUR"
func (v Value) Abs() Value {
	return Value{amount: v.amount.Abs(), currency: v.currency}
}

// Neg returns the negative value.
//
//	MustFromString("-123.456 ISO4217-EUR").Neg() // Returns "123.456 ISO4217-EUR"
//	MustFromString("123.456 ISO4217-EUR").Neg()  // Returns "-123.456 ISO4217-EUR"
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

// SplitWithSmallestUnit returns the value of v split into a list of n values.
// If the value can't be split evenly, the remainder will be distributed round-robin amongst the parts.
// The resulting values will always be multiple of smallestUnit, if that's not possible an error will be returned.
//
//	MustFromString("-11.11").SplitWithSmallestUnit(3, MustFromString("0.01"))                         // Returns the three values `-3.71`, `-3.7`, `-3.7`.
//	MustFromString("-11.11 ISO4217-EUR").SplitWithSmallestUnit(3, MustFromString("0.01 ISO4217-EUR")) // Returns the three EUR values `-3.71`, `-3.7`, `-3.7`.
//	MustFromString("-11.11").SplitWithSmallestUnit(3, MustFromString("0.1"))                          // Returns an error, as the value can't be split into parts that are multiple of the smallest unit (0.1).
func (v Value) SplitWithSmallestUnit(n int, smallestUnit Value) ([]Value, error) {
	if n <= 0 {
		return nil, fmt.Errorf("number of parts must not be negative")
	}
	if smallestUnit.Sign() <= 0 {
		return nil, fmt.Errorf("smallest unit with %s is outside the allowed range", smallestUnit)
	}
	if smallestUnit.currency != v.currency {
		return nil, &ErrorDifferentCurrencies{v.currency, smallestUnit.currency}
	}

	// Negate smallest unit if the value is negative.
	// This way we will always have positive amounts of smallest units.
	smallestUnitSigned := smallestUnit.Decimal()
	if v.IsNegative() {
		smallestUnitSigned = smallestUnitSigned.Neg()
	}

	// Get amount of smallest units that have to be distributed.
	q, r := v.amount.QuoRem(smallestUnitSigned, 0)
	if !r.IsZero() {
		return nil, fmt.Errorf("value is not a multiple of the smallest unit %s", smallestUnit)
	}

	// Get amount of smallest units per part.
	qPart, rPart := q.QuoRem(decimal.NewFromInt(int64(n)), 0)
	if !rPart.IsInteger() {
		// This shouldn't happen, n and q are supposed to be integers.
		panic("The remainder must be an integer")
	}
	// The first largerParts number of result values contain one more smallest unit.
	// largerParts can't be larger than n, which is an int.
	largerParts := int(rPart.IntPart())

	value1 := Value{amount: qPart.Mul(smallestUnitSigned), currency: v.currency}
	value2 := Value{amount: value1.amount.Add(smallestUnitSigned), currency: v.currency}

	// Build list of values.
	values := make([]Value, n)
	for i := range values {
		if i >= largerParts {
			values[i] = value1 // qPart * smallestUnit
		} else {
			values[i] = value2 // qPart * smallestUnit + smallestUnit
		}
	}

	return values, nil
}

// SplitWithDecimals returns the value of v split into a list of n values.
// If the value can't be split evenly, the remainder will be distributed round-robin amongst the parts.
// The smallest unit that the value is split into is calculated by 10^(-decimalPlaces).
//
//	MustFromString("-11.11 ISO4217-EUR").SplitWithDecimals(3, 2) // Returns the three EUR values `-3.71`, `-3.7`, `-3.7`.
//	MustFromString("-11.11 ISO4217-EUR").SplitWithDecimals(3, 1) // Returns an error, as the value can't be split into parts that are multiple of the smallest unit (0.1).
func (v Value) SplitWithDecimals(n int, decimalPlaces int) ([]Value, error) {
	if decimalPlaces <= math.MinInt32 || decimalPlaces > math.MaxInt32 { // Exclude MinInt32 from valid range, as we want to invert the number.
		return nil, fmt.Errorf("decimal places (%d) is outside the allowed range", decimalPlaces)
	}

	smallestUnit := FromDecimal(decimal.New(1, -int32(decimalPlaces)), v.currency)

	return v.SplitWithSmallestUnit(n, smallestUnit)
}

// Split returns the value of v split into a list of n values.
// If the value can't be split evenly, the remainder will be distributed round-robin amongst the parts.
// The smallest unit is determined by the currency of the given value.
//
//	MustFromString("-11.11 ISO4217-EUR").Split(3) // Returns the three EUR values `-3.71`, `-3.7`, `-3.7`.
//	MustFromString("-11.11 ISO4217-VND").Split(3) // Returns an error, as the value can't be split into parts that are multiple of the smallest unit (1).
//	MustFromString("-1111 ISO4217-VND").Split(3)  // Returns the three VND values `-371`, `-370`, `-370`.
//	MustFromString("-11.11").Split(3)             // Returns an error, as there is no smallest unit.
func (v Value) Split(n int) ([]Value, error) {
	var smallestUnit Value // Default value is a smallest unit of 0. Which means that there is no smallest unit.
	if v.currency != nil {
		smallestUnit = v.currency.SmallestUnit()
	}

	return v.SplitWithSmallestUnit(n, smallestUnit)
}
