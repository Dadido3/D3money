// Copyright (c) 2021-2022 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"golang.org/x/text/language"
)

// testCurrency implements a custom currency for testing purposes.
type testCurrency struct {
	name, standard, code, symbol, narrowSymbol string
	numericCode                                int
	uniqueID                                   int32
	smallestUnit                               decimal.Decimal
	noSmallestUnitCurrency                     bool // If true, SmallestUnit will return a Value with no currency set.
}

func (c *testCurrency) Name() string                          { return c.code }
func (c *testCurrency) Standard() string                      { return c.standard }
func (c *testCurrency) UniqueID() int32                       { return c.uniqueID }
func (c *testCurrency) UniqueCode() string                    { return c.Standard() + "-" + c.Code() }
func (c *testCurrency) NumericCode() int                      { return c.numericCode }
func (c *testCurrency) Code() string                          { return c.code }
func (c *testCurrency) Symbol(lang language.Tag) string       { return c.symbol }
func (c *testCurrency) NarrowSymbol(lang language.Tag) string { return c.narrowSymbol }
func (c *testCurrency) SmallestUnit() Value {
	if c.noSmallestUnitCurrency {
		return Value{c.smallestUnit, nil}
	} else {
		return Value{c.smallestUnit, c}
	}
}

var (
	testCurrency1          Currency = &testCurrency{"Bar", "FOO", "BAR", "|", "|", 1, 1, decimal.New(1, -2), false}
	testCurrencyCollision1 Currency = &testCurrency{"Bar", "FOZ", "BAR", "‖", "‖", 2, 2, decimal.New(1, -2), false} // Collides with testCurrency1 on its code.
	testCurrencyCollision2 Currency = &testCurrency{"Bar", "FOO", "BAR", "⫼", "⫼", 3, 3, decimal.New(1, -2), false} // Collides with testCurrency1 on its code and unique code.
	testCurrencyCollision3 Currency = &testCurrency{"Baz", "FOZ", "BAZ", "⟊", "⟊", 4, 1, decimal.New(1, -2), false} // Collides with testCurrency1 on its unique ID.
	testCurrencyCollision4 Currency = &testCurrency{"Baz", "FOZ", "BAZ", "˥", "˥", 1, 5, decimal.New(1, -2), false} // Collides with testCurrency1 on its numeric code.

	testCurrencyIllegal1 Currency = &testCurrency{"Bar", "FOO", "BAR", "|", "|", 1, 0, decimal.New(1, -2), false} // Variant of testCurrency1 that contains an illegal unique ID.
	testCurrencyIllegal2 Currency = &testCurrency{"Bar", "FOO", "Bar", "|", "|", 1, 1, decimal.New(1, -2), false} // Variant of testCurrency1 that contains an illegal code.
	testCurrencyIllegal3 Currency = &testCurrency{"Bar", "Foo", "BAR", "|", "|", 1, 1, decimal.New(1, -2), false} // Variant of testCurrency1 that contains an illegal standard string.
	testCurrencyIllegal4 Currency = &testCurrency{"Bar", "FOO", "BAR", "|", "|", 1, 1, decimal.New(-1, 0), false} // Variant of testCurrency1 that contains an illegal smallest unit value.
	testCurrencyIllegal5 Currency = &testCurrency{"Bar", "FOO", "BAR", "|", "|", 1, 1, decimal.New(1, -1), true}  // Variant of testCurrency1 that contains an illegal smallest unit value.
	testCurrencyIllegal6 Currency = &testCurrency{"Bar", "FOO", "BAR", "|", "", 1, 1, decimal.New(1, -2), false}  // Variant of testCurrency1 that contains an illegal symbol combination.
	testCurrencyIllegal7 Currency = &testCurrency{"Bar", "FOO", "BAR", "", "|", 1, 1, decimal.New(1, -2), false}  // Variant of testCurrency1 that contains an illegal symbol combination.
)

func TestFooBarCurrencies(t *testing.T) {
	currencies := []Currency{
		testCurrency1,
		testCurrencyCollision1,
		testCurrencyCollision2,
		testCurrencyCollision3,
		testCurrencyCollision4,
	}

	for _, currency := range currencies {
		if err := ValidateCurrency(currency); err != nil {
			t.Errorf("ValidateCurrency() failed to validate currency %s: %v", helperCurrencyUniqueCode(currency), err)
		}
	}
}

func TestValidateCurrency(t *testing.T) {
	currencies := []Currency{
		testCurrencyIllegal1,
		testCurrencyIllegal2,
		testCurrencyIllegal3,
		testCurrencyIllegal4,
		testCurrencyIllegal5,
		testCurrencyIllegal6,
		testCurrencyIllegal7,
	}

	for _, currency := range currencies {
		if err := ValidateCurrency(currency); err == nil {
			t.Errorf("ValidateCurrency() failed to validate currency %s. Expected error", helperCurrencyUniqueCode(currency))
		}
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    Value
		wantErr bool
	}{
		{"1", args{"-10000.123"}, Value{decimal.New(-10000123, -3), nil}, false},
		{"2", args{"-10000.123 ISO4217-EUR"}, Value{decimal.New(-10000123, -3), ISO4217Currencies.ByCode("EUR")}, false},
		{"3", args{"-10000.123 EUR"}, Value{decimal.Decimal{}, nil}, true},
		{"4", args{"-10000.123 FOO-BAR"}, Value{decimal.Decimal{}, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromStringAndCurrency(t *testing.T) {
	type args struct {
		str string
		cur Currency
	}
	tests := []struct {
		name    string
		args    args
		want    Value
		wantErr bool
	}{
		{"1", args{"-10000.123", nil}, Value{decimal.New(-10000123, -3), nil}, false},
		{"2", args{"-10000.123 ISO4217-EUR", nil}, Value{decimal.Decimal{}, nil}, true},
		{"3", args{"-10000.123", ISO4217Currencies.ByCode("EUR")}, Value{decimal.New(-10000123, -3), ISO4217Currencies.ByCode("EUR")}, false},
		{"4", args{"-10000.123 ISO4217-EUR", ISO4217Currencies.ByCode("EUR")}, Value{decimal.New(-10000123, -3), ISO4217Currencies.ByCode("EUR")}, false},
		{"5", args{"-10000.123 ISO4217-USD", ISO4217Currencies.ByCode("EUR")}, Value{decimal.Decimal{}, nil}, true},
		{"6", args{"-10000.123 FOO-BAR", nil}, Value{decimal.Decimal{}, nil}, true},
		{"7", args{"-10000.123", testCurrency1}, Value{decimal.New(-10000123, -3), testCurrency1}, false},
		{"8", args{"-10000.123 FOO-BAR", testCurrency1}, Value{decimal.New(-10000123, -3), testCurrency1}, false},
		{"9", args{"-10000.123 ISO4217-EUR", testCurrency1}, Value{decimal.Decimal{}, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromStringAndCurrency(tt.args.str, tt.args.cur)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromStringAndCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FromStringAndCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
