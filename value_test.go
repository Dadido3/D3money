// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

// testCurrency implements a custom currency for testing purposes.
type testCurrency struct {
	name, standard, code, symbol, shortSymbol string
	uniqueID                                  int32
	decimalPlaces                             int
	hasSmallestUnit                           bool
}

func (c *testCurrency) Name() string               { return c.code }
func (c *testCurrency) StandardName() string       { return c.standard }
func (c *testCurrency) UniqueID() int32            { return c.uniqueID }
func (c *testCurrency) UniqueCode() string         { return c.StandardName() + "-" + c.Code() }
func (c *testCurrency) Code() string               { return c.code }
func (c *testCurrency) Symbol() string             { return c.symbol }
func (c *testCurrency) ShortSymbol() string        { return c.shortSymbol }
func (c *testCurrency) DecimalPlaces() (int, bool) { return c.decimalPlaces, c.hasSmallestUnit }

var (
	testCurrency1          Currency = &testCurrency{"Bar", "FOO", "BAR", "|", "|", 1, 2, true}
	testCurrencyCollision1 Currency = &testCurrency{"Bar", "FOZ", "BAR", "‖", "‖", 2, 2, true} // Collides with testCurrency1 on its code.
	testCurrencyCollision2 Currency = &testCurrency{"Bar", "FOO", "BAR", "⫼", "⫼", 3, 2, true} // Collides with testCurrency1 on its code and unique code.
	testCurrencyCollision3 Currency = &testCurrency{"Baz", "FOZ", "BAZ", "⟊", "⟊", 1, 2, true} // Collides with testCurrency1 on its unique ID.
)

func TestFooBarCurrency(t *testing.T) {
	for _, currency := range []Currency{testCurrency1, testCurrencyCollision1, testCurrencyCollision2, testCurrencyCollision3} {
		if err := ValidateCurrency(currency); err != nil {
			t.Errorf("Failed to validate currency %q: %v", currency, err)
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
			if !reflect.DeepEqual(got, tt.want) {
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStringAndCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
