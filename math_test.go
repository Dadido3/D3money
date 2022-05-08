// Copyright (c) 2021-2022 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		first  Value
		values []Value
	}
	tests := []struct {
		name    string
		args    args
		want    Value
		wantErr bool
	}{
		{"comment_1", args{MustFromString("12.34 ISO4217-EUR"), []Value{MustFromString("12.34 ISO4217-EUR")}}, MustFromString("24.68 ISO4217-EUR"), false},
		{"comment_2", args{MustFromString("12.34 ISO4217-EUR"), []Value{MustFromString("12.34")}}, Value{}, true},
		{"1", args{MustFromString("12.34 ISO4217-EUR"), nil}, MustFromString("12.34 ISO4217-EUR"), false},
		{"2", args{MustFromString("-12.34"), []Value{MustFromString("12.34")}}, MustFromString("0"), false},
		{"3", args{MustFromString("-12.34"), []Value{MustFromString("12.34"), MustFromString("12.34")}}, MustFromString("12.34"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.args.first, tt.args.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if equal, _ := got.Equal(tt.want); !equal {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
