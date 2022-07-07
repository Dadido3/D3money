// Copyright (c) 2022 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"golang.org/x/text/language"
)

type ValueFormatterDisplay int

const (
	DisplaySymbol     ValueFormatterDisplay = iota // Formatter will use the currency symbol. E.g. "US$". This is the default.
	DisplayNarrow                                  // Formatter will use the narrow currency symbol. E.g. "$".
	DisplayCode                                    // Formatter will use the currency code. E.g. "USD".
	DisplayUniqueCode                              // Formatter will use the currency's unique code. E.g. "ISO4217-USD".
	DisplayName                                    // Formatter will use the currency name. E.g. "US dollar" or "US dollars".
)

// ValueFormatter encapsulates a value and provides formatting capabilities.
type ValueFormatter struct {
	value    Value
	language language.Tag
	display  ValueFormatterDisplay
}

/*func (v ValueFormatter) Format(f fmt.State, verb rune) {
	switch verb {
	case 's', 'q':
		return "Hey"
	}

	return "Ho"
}*/

func (v ValueFormatter) Display(display ValueFormatterDisplay) ValueFormatter {
	v.display = display
	return v
}

func (v ValueFormatter) String() string {
	return ""
}
