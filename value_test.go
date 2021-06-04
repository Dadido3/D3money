package d3money

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

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
		{"1", args{"-10000.123"}, Value{decimal.New(-10000123, -3), nil, nil}, false},
		{"2", args{"-10000.123 ISO4217-EUR"}, Value{decimal.Decimal{}, nil, nil}, true},
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
		{"1", args{"-10000.123", nil}, Value{decimal.New(-10000123, -3), nil, nil}, false},
		{"2", args{"-10000.123 ISO4217-EUR", nil}, Value{decimal.Decimal{}, nil, nil}, true},
		{"3", args{"-10000.123", ISO4217Currencies["EUR"]}, Value{decimal.New(-10000123, -3), ISO4217Currencies["EUR"], nil}, false},
		{"4", args{"-10000.123 ISO4217-EUR", ISO4217Currencies["EUR"]}, Value{decimal.New(-10000123, -3), ISO4217Currencies["EUR"], nil}, false},
		{"5", args{"-10000.123 ISO4217-USD", ISO4217Currencies["EUR"]}, Value{decimal.Decimal{}, nil, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromStringWithCurrency(tt.args.str, tt.args.cur)
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

func TestNewFromStringWithCurrencies(t *testing.T) {
	type args struct {
		str        string
		cur        Currency
		currencies map[string]Currency
	}
	tests := []struct {
		name    string
		args    args
		want    Value
		wantErr bool
	}{
		{"1", args{"-10000.123", nil, nil}, Value{decimal.New(-10000123, -3), nil, nil}, false},
		{"2", args{"-10000.123", nil, ISO4217Currencies}, Value{decimal.New(-10000123, -3), nil, nil}, false},
		{"3", args{"-10000.123 ISO4217-EUR", nil, ISO4217Currencies}, Value{decimal.New(-10000123, -3), ISO4217Currencies["EUR"], nil}, false},
		{"4", args{"-10000.123", ISO4217Currencies["EUR"], ISO4217Currencies}, Value{decimal.New(-10000123, -3), ISO4217Currencies["EUR"], nil}, false},
		{"5", args{"-10000.123 ISO4217-EUR", ISO4217Currencies["EUR"], ISO4217Currencies}, Value{decimal.New(-10000123, -3), ISO4217Currencies["EUR"], nil}, false},
		{"6", args{"-10000.123", ISO4217Currencies["EUR"], nil}, Value{decimal.New(-10000123, -3), ISO4217Currencies["EUR"], nil}, false},
		{"7", args{"-10000.123 ISO4217-EUR", ISO4217Currencies["EUR"], nil}, Value{decimal.New(-10000123, -3), ISO4217Currencies["EUR"], nil}, false},
		{"8", args{"-10000.123 ISO4217-EUR", nil, nil}, Value{decimal.Decimal{}, nil, nil}, true},
		{"9", args{"-10000.123 ISO4217-USD", ISO4217Currencies["EUR"], ISO4217Currencies}, Value{decimal.Decimal{}, nil, nil}, true},
		{"10", args{"-10000.123 ", nil, nil}, Value{decimal.Decimal{}, nil, nil}, true},
		{"11", args{"-10000.123  ISO4217-EUR", nil, nil}, Value{decimal.Decimal{}, nil, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromStringWithCurrencies(tt.args.str, tt.args.cur, tt.args.currencies)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromStringWithCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromStringWithCurrencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test1(t *testing.T) {
	val, err := NewFromStringWithCurrencies("-12345.67 ISO4217-EUR", nil, ISO4217Currencies)
	if err != nil {
		t.Errorf("NewFromStringWithCurrencies() failed: %v", err)
	}

	if str := val.String(); str != "-12345.67\u00A0ISO4217-EUR" {
		t.Errorf("val.String() = %q: want %q", str, "-12345.67\u00A0ISO4217-EUR")
	}

	val2, err := NewFromStringWithCurrencies("-12345.67", nil, ISO4217Currencies)
	if err != nil {
		t.Errorf("NewFromStringWithCurrencies() failed: %v", err)
	}

	if str := val2.String(); str != "-12345.67" {
		t.Errorf("val.String() = %q: want %q", str, "-12345.67")
	}
}
