// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"testing"
)

func TestNewCurrencyCollection(t *testing.T) {

	// No error should be returned for currencies with colliding codes.
	if _, err := NewCurrencyCollection("Test1", "", []Currency{testCurrency1, testCurrencyCollision1}); err != nil {
		t.Errorf("NewCurrencyCollection() error = %v", err)
	}

	// Error should be returned when trying to add currency of a different standard.
	if _, err := NewCurrencyCollection("Test1", "FOO", []Currency{testCurrency1, testCurrencyCollision1}); err == nil {
		t.Error("NewCurrencyCollection() with different standard did not fail")
	}

	// Error should be returned for currencies with colliding codes.
	if _, err := NewCurrencyCollection("Test1", "FOO", []Currency{testCurrency1, testCurrencyCollision2}); err == nil {
		t.Error("NewCurrencyCollection() with colliding codes did not fail")
	}

}
