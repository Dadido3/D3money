package d3money

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

// fooBarCurrency implements a custom currency for testing purposes.
type fooBarCurrency struct{}

func (c *fooBarCurrency) Name() string               { return "Bar" }
func (c *fooBarCurrency) StandardName() string       { return "FOO" }
func (c *fooBarCurrency) UniqueID() int32            { return -1 }
func (c *fooBarCurrency) UniqueCode() string         { return "FOO-BAR" }
func (c *fooBarCurrency) Code() string               { return "BAR" }
func (c *fooBarCurrency) Symbol() string             { return "❚" }
func (c *fooBarCurrency) ShortSymbol() string        { return "❚" }
func (c *fooBarCurrency) DecimalPlaces() (int, bool) { return 2, true }

var FooBarCurrency Currency = new(fooBarCurrency)

func TestNewFromString(t *testing.T) {
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
		{"2", args{"-10000.123 ISO4217-EUR"}, Value{decimal.New(-10000123, -3), ISO4217Currencies.CurrencyByCode("EUR")}, false},
		{"3", args{"-10000.123 EUR"}, Value{decimal.Decimal{}, nil}, true},
		{"4", args{"-10000.123 FOO-BAR"}, Value{decimal.Decimal{}, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromStringWithCurrency(t *testing.T) {
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
		{"3", args{"-10000.123", ISO4217Currencies.CurrencyByCode("EUR")}, Value{decimal.New(-10000123, -3), ISO4217Currencies.CurrencyByCode("EUR")}, false},
		{"4", args{"-10000.123 ISO4217-EUR", ISO4217Currencies.CurrencyByCode("EUR")}, Value{decimal.New(-10000123, -3), ISO4217Currencies.CurrencyByCode("EUR")}, false},
		{"5", args{"-10000.123 ISO4217-USD", ISO4217Currencies.CurrencyByCode("EUR")}, Value{decimal.Decimal{}, nil}, true},
		{"6", args{"-10000.123 FOO-BAR", nil}, Value{decimal.Decimal{}, nil}, true},
		{"7", args{"-10000.123", FooBarCurrency}, Value{decimal.New(-10000123, -3), FooBarCurrency}, false},
		{"8", args{"-10000.123 FOO-BAR", FooBarCurrency}, Value{decimal.New(-10000123, -3), FooBarCurrency}, false},
		{"9", args{"-10000.123 ISO4217-EUR", FooBarCurrency}, Value{decimal.Decimal{}, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromStringAndCurrency(tt.args.str, tt.args.cur)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromStringWithCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromStringWithCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test1(t *testing.T) {
	/*val, err := NewFromStringWithCurrencies("-12345.67 ISO4217-EUR", nil, iso4217Currencies)
	if err != nil {
		t.Errorf("NewFromStringWithCurrencies() failed: %v", err)
	}

	if str := val.String(); str != "-12345.67\u00A0ISO4217-EUR" {
		t.Errorf("val.String() = %q: want %q", str, "-12345.67\u00A0ISO4217-EUR")
	}

	val2, err := NewFromStringWithCurrencies("-12345.67", nil, iso4217Currencies)
	if err != nil {
		t.Errorf("NewFromStringWithCurrencies() failed: %v", err)
	}

	if str := val2.String(); str != "-12345.67" {
		t.Errorf("val.String() = %q: want %q", str, "-12345.67")
	}*/
}
