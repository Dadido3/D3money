package d3money

import "testing"

func TestCurrencies(t *testing.T) {
	for _, currency := range Currencies.All() {
		if err := ValidateCurrency(currency); err != nil {
			t.Errorf("Failed to validate currency %s: %v", currency, err)
		}
	}
}
