// Copyright (c) 2021-2022 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import "testing"

func TestCurrencies(t *testing.T) {
	for _, currency := range Currencies.All() {
		if err := ValidateCurrency(currency); err != nil {
			t.Errorf("Failed to validate currency %s: %v", helperCurrencyUniqueCode(currency), err)
		}
	}
}
