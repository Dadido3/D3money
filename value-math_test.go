// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"testing"
)

func TestValue_Equal(t *testing.T) {
	type args struct {
		comp Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    bool
		wantErr bool
	}{
		{"1", MustFromString("-1234567.89"), args{MustFromString("-1234567.89")}, true, false},
		{"2", MustFromString("-1234567.89"), args{MustFromString("-1234567.88")}, false, false},
		{"3", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89 ISO4217-EUR")}, true, false},
		{"4", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.88 ISO4217-EUR")}, false, false},
		{"5", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89")}, false, true},
		{"6", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.88")}, false, true},
		{"7", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89 ISO4217-USD")}, false, true},
		{"8", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.88 ISO4217-USD")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Equal(tt.args.comp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.Equal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_GreaterThan(t *testing.T) {
	type args struct {
		comp Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    bool
		wantErr bool
	}{
		{"1", MustFromString("-1234567.89"), args{MustFromString("-1234567.90")}, true, false},
		{"2", MustFromString("-1234567.89"), args{MustFromString("-1234567.89")}, false, false},
		{"3", MustFromString("-1234567.89"), args{MustFromString("-1234567.88")}, false, false},
		{"4", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89 ISO4217-USD")}, false, true},
		{"5", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.GreaterThan(tt.args.comp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.GreaterThan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value.GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_GreaterThanOrEqual(t *testing.T) {
	type args struct {
		comp Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    bool
		wantErr bool
	}{
		{"1", MustFromString("-1234567.89"), args{MustFromString("-1234567.90")}, true, false},
		{"2", MustFromString("-1234567.89"), args{MustFromString("-1234567.89")}, true, false},
		{"3", MustFromString("-1234567.89"), args{MustFromString("-1234567.88")}, false, false},
		{"4", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89 ISO4217-USD")}, false, true},
		{"5", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.GreaterThanOrEqual(tt.args.comp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.GreaterThanOrEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value.GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_LessThan(t *testing.T) {
	type args struct {
		comp Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    bool
		wantErr bool
	}{
		{"1", MustFromString("-1234567.89"), args{MustFromString("-1234567.90")}, false, false},
		{"2", MustFromString("-1234567.89"), args{MustFromString("-1234567.89")}, false, false},
		{"3", MustFromString("-1234567.89"), args{MustFromString("-1234567.88")}, true, false},
		{"4", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89 ISO4217-USD")}, false, true},
		{"5", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.LessThan(tt.args.comp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.LessThan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value.LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_LessThanOrEqual(t *testing.T) {
	type args struct {
		comp Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    bool
		wantErr bool
	}{
		{"1", MustFromString("-1234567.89"), args{MustFromString("-1234567.90")}, false, false},
		{"2", MustFromString("-1234567.89"), args{MustFromString("-1234567.89")}, true, false},
		{"3", MustFromString("-1234567.89"), args{MustFromString("-1234567.88")}, true, false},
		{"4", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89 ISO4217-USD")}, false, true},
		{"5", MustFromString("-1234567.89 ISO4217-EUR"), args{MustFromString("-1234567.89")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.LessThanOrEqual(tt.args.comp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.LessThanOrEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value.LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Float64(t *testing.T) {
	tests := []struct {
		name      string
		v         Value
		wantF     float64
		wantExact bool
	}{
		{"1", MustFromString("2"), 2, true},
		{"2", MustFromString("1.0000000000000002"), 1.0000000000000002, false},
		{"3", MustFromString("1"), 1, true},
		{"4", MustFromString("1 ISO4217-EUR"), 1, true},
		{"5", MustFromString("0.01171875"), 0.01171875, true},
		{"6", MustFromString("0"), 0, true},
		{"7", MustFromString("-1"), -1, true},
		{"8", MustFromString("-2"), -2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, gotExact := tt.v.Float64()
			if gotF != tt.wantF {
				t.Errorf("Value.Float64() gotF = %v, want %v", gotF, tt.wantF)
			}
			if gotExact != tt.wantExact {
				t.Errorf("Value.Float64() gotExact = %v, want %v", gotExact, tt.wantExact)
			}
		})
	}
}

func TestValue_InexactFloat64(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		want float64
	}{
		{"1", MustFromString("2"), 2},
		{"2", MustFromString("1.0000000000000002"), 1.0000000000000002},
		{"3", MustFromString("1"), 1},
		{"4", MustFromString("1 ISO4217-EUR"), 1},
		{"5", MustFromString("0.01171875"), 0.01171875},
		{"6", MustFromString("0"), 0},
		{"7", MustFromString("-1"), -1},
		{"8", MustFromString("-2"), -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.InexactFloat64(); got != tt.want {
				t.Errorf("Value.InexactFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Add(t *testing.T) {
	type args struct {
		v2 Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    Value
		wantErr bool
	}{
		{"1", MustFromString("0"), args{MustFromString("-1")}, MustFromString("-1"), false},
		{"2", MustFromString("1 ISO4217-EUR"), args{MustFromString("-1 ISO4217-EUR")}, MustFromString("0 ISO4217-EUR"), false},
		{"3", MustFromString("1"), args{MustFromString("-0.0000000000000000000000000000001")}, MustFromString("0.9999999999999999999999999999999"), false},
		{"4", MustFromString("0"), args{MustFromString("0 ISO4217-EUR")}, Value{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Add(tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if equal, _ := got.Equal(tt.want); !equal {
				t.Errorf("Value.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Sub(t *testing.T) {
	type args struct {
		v2 Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    Value
		wantErr bool
	}{
		{"1", MustFromString("0"), args{MustFromString("-1")}, MustFromString("1"), false},
		{"2", MustFromString("1 ISO4217-EUR"), args{MustFromString("-1 ISO4217-EUR")}, MustFromString("2 ISO4217-EUR"), false},
		{"3", MustFromString("1"), args{MustFromString("-0.0000000000000000000000000000001")}, MustFromString("1.0000000000000000000000000000001"), false},
		{"4", MustFromString("0"), args{MustFromString("0 ISO4217-EUR")}, Value{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Sub(tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if equal, _ := got.Equal(tt.want); !equal {
				t.Errorf("Value.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Mul(t *testing.T) {
	type args struct {
		v2 Value
	}
	tests := []struct {
		name    string
		v       Value
		args    args
		want    Value
		wantErr bool
	}{
		{"1", MustFromString("0"), args{MustFromString("-1")}, MustFromString("0"), false},
		{"2", MustFromString("1 ISO4217-EUR"), args{MustFromString("-1")}, MustFromString("-1 ISO4217-EUR"), false},
		{"3", MustFromString("1"), args{MustFromString("-1 ISO4217-EUR")}, MustFromString("-1 ISO4217-EUR"), false},
		{"4", MustFromString("1"), args{MustFromString("-0.0000000000000000000000000000001")}, MustFromString("-0.0000000000000000000000000000001"), false},
		{"5", MustFromString("0 ISO4217-EUR"), args{MustFromString("0 ISO4217-EUR")}, Value{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Mul(tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.Mul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if equal, _ := got.Equal(tt.want); !equal {
				t.Errorf("Value.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Abs(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		want Value
	}{
		{"1", MustFromString("0"), MustFromString("0")},
		{"2", MustFromString("-1.234"), MustFromString("1.234")},
		{"3", MustFromString("1.234"), MustFromString("1.234")},
		{"4", MustFromString("1.234 ISO4217-EUR"), MustFromString("1.234 ISO4217-EUR")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Abs()
			if equal, _ := got.Equal(tt.want); !equal {
				t.Errorf("Value.Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Neg(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		want Value
	}{
		{"1", MustFromString("0"), MustFromString("0")},
		{"2", MustFromString("-1.234"), MustFromString("1.234")},
		{"3", MustFromString("1.234"), MustFromString("-1.234")},
		{"4", MustFromString("1.234 ISO4217-EUR"), MustFromString("-1.234 ISO4217-EUR")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Neg()
			if equal, _ := got.Equal(tt.want); !equal {
				t.Errorf("Value.Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_Sign(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		want int
	}{
		{"1", MustFromString("0"), 0},
		{"2", MustFromString("-1.234"), -1},
		{"3", MustFromString("1.234"), 1},
		{"4", MustFromString("1.234 ISO4217-EUR"), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sign(); got != tt.want {
				t.Errorf("Value.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_IsPositive(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		want bool
	}{
		{"1", MustFromString("0"), false},
		{"2", MustFromString("-1.234"), false},
		{"3", MustFromString("1.234"), true},
		{"4", MustFromString("1.234 ISO4217-EUR"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsPositive(); got != tt.want {
				t.Errorf("Value.IsPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_IsNegative(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		want bool
	}{
		{"1", MustFromString("0"), false},
		{"2", MustFromString("-1.234"), true},
		{"3", MustFromString("1.234"), false},
		{"4", MustFromString("1.234 ISO4217-EUR"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsNegative(); got != tt.want {
				t.Errorf("Value.IsNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_IsZero(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		want bool
	}{
		{"1", MustFromString("0"), true},
		{"2", MustFromString("-1.234"), false},
		{"3", MustFromString("1.234"), false},
		{"4", MustFromString("1.234 ISO4217-EUR"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsZero(); got != tt.want {
				t.Errorf("Value.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}