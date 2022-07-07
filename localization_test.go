// Copyright (c) 2022 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestCurrency_Symbol_NarrowSymbol(t *testing.T) {
	type args struct {
		langTag string
	}
	tests := []struct {
		c                            Currency
		args                         args
		wantSymbol, wantNarrowSymbol string
	}{
		{ISO4217Currencies.ByCode("EUR"), args{"und"}, "EUR", "EUR"},
		{ISO4217Currencies.ByCode("EUR"), args{"de"}, "€", "€"},
		{ISO4217Currencies.ByCode("EUR"), args{"en"}, "€", "€"},
		{ISO4217Currencies.ByCode("EUR"), args{"en-AU"}, "€", "€"},
		{ISO4217Currencies.ByCode("EUR"), args{"en-GB"}, "€", "€"},
		{ISO4217Currencies.ByCode("EUR"), args{"en-US"}, "€", "€"},

		{ISO4217Currencies.ByCode("USD"), args{"und"}, "USD", "USD"},
		{ISO4217Currencies.ByCode("USD"), args{"de"}, "$", "$"},
		{ISO4217Currencies.ByCode("USD"), args{"en"}, "$", "$"},
		{ISO4217Currencies.ByCode("USD"), args{"en-AU"}, "USD", "$"},
		{ISO4217Currencies.ByCode("USD"), args{"en-GB"}, "US$", "$"},
		{ISO4217Currencies.ByCode("USD"), args{"en-US"}, "$", "$"},

		{ISO4217Currencies.ByCode("INR"), args{"und"}, "INR", "INR"},
		{ISO4217Currencies.ByCode("INR"), args{"de"}, "₹", "₹"},
		{ISO4217Currencies.ByCode("INR"), args{"en"}, "₹", "₹"},
		{ISO4217Currencies.ByCode("INR"), args{"en-AU"}, "INR", "₹"},
		{ISO4217Currencies.ByCode("INR"), args{"en-GB"}, "₹", "₹"},
		{ISO4217Currencies.ByCode("INR"), args{"en-US"}, "₹", "₹"},

		{ISO4217Currencies.ByCode("AUD"), args{"und"}, "AUD", "AUD"},
		{ISO4217Currencies.ByCode("AUD"), args{"de"}, "AU$", "$"},
		{ISO4217Currencies.ByCode("AUD"), args{"en"}, "A$", "$"},
		{ISO4217Currencies.ByCode("AUD"), args{"en-AU"}, "$", "$"},
		{ISO4217Currencies.ByCode("AUD"), args{"en-GB"}, "A$", "$"},
		{ISO4217Currencies.ByCode("AUD"), args{"en-US"}, "A$", "$"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s in %s", tt.c.UniqueCode(), tt.args.langTag)
		t.Run(name, func(t *testing.T) {
			tag, err := language.Parse(tt.args.langTag)
			if err != nil {
				t.Fatalf("Couldn't parse locale tag %q", tt.args.langTag)
			}
			if got := tt.c.Symbol(tag); got != tt.wantSymbol {
				t.Errorf("ISO4217Currency.Symbol() = %v, want %v. Locale: %q", got, tt.wantSymbol, tag)
			}
			if got := tt.c.NarrowSymbol(tag); got != tt.wantNarrowSymbol {
				t.Errorf("ISO4217Currency.NarrowSymbol() = %v, want %v. Locale: %q", got, tt.wantNarrowSymbol, tag)
			}
		})
	}
}
