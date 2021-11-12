// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import "fmt"

// ErrorInvalidCurrency is returned when the definition of a currency contains invalid or illegal values.
type ErrorInvalidCurrency struct{ msg string }

func (e *ErrorInvalidCurrency) Error() string { return e.msg }

// ErrorDifferentCurrencies is returned when the currencies between two values don't match.
type ErrorDifferentCurrencies struct{ c1, c2 Currency }

func (e *ErrorDifferentCurrencies) Error() string {
	if e.c1 == nil {
		return fmt.Sprintf("the monetary values have two different currencies: no currency and %q", e.c2)
	}
	if e.c2 == nil {
		return fmt.Sprintf("the monetary values have two different currencies: %q and no currency", e.c1)
	}

	return fmt.Sprintf("the monetary values have two different currencies: %q and %q", e.c1, e.c2)
}

// ErrorCantFindUniqueID is returned when no currency can be found for a given unique ID.
type ErrorCantFindUniqueID struct{ uniqueID int32 }

func (e *ErrorCantFindUniqueID) Error() string {
	return fmt.Sprintf("can't find currency with unique ID %d", e.uniqueID)
}

// ErrorCantFindUniqueCode is returned when no currency can be found for a given unique code.
type ErrorCantFindUniqueCode struct{ uniqueCode string }

func (e *ErrorCantFindUniqueCode) Error() string {
	return fmt.Sprintf("can't find currency with unique ID %q", e.uniqueCode)
}
