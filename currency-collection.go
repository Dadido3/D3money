// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"fmt"
	"sync"
)

// CurrencyCollection is a container for currencies.
type CurrencyCollection struct {
	mutex sync.RWMutex

	name string

	// Name of the currency standard that this collection is limited to.
	// If set != "", no other currency standard other than the defined one will be allowed.
	currencyStandard string

	all           []Currency
	hasCurrency   map[Currency]struct{}
	byUniqueID    map[int32]Currency
	byUniqueCode  map[string]Currency
	byNumericCode map[int]Currency
	byCode        map[string]Currency
}

// NewCurrencyCollection takes one or more lists of currencies and returns them as a collection.
//
// With currencyStandard you can limit the collection to only allow currencies of a specific standard.
// Set it to "" if you want to allow multiple currency standards.
func NewCurrencyCollection(name string, currencyStandard string, listsOfCurrencies ...[]Currency) (*CurrencyCollection, error) {
	cc := &CurrencyCollection{
		name:             name,
		currencyStandard: currencyStandard,
		hasCurrency:      map[Currency]struct{}{},
		byUniqueID:       map[int32]Currency{},
		byUniqueCode:     map[string]Currency{},
		byNumericCode:    nil,
		byCode:           nil,
	}

	// Enable search by code and numeric code if the collection is set to a single standard.
	if cc.currencyStandard != "" {
		cc.byNumericCode, cc.byCode = map[int]Currency{}, map[string]Currency{}
	}

	for _, listOfCurrencies := range listsOfCurrencies {
		if err := cc.Add(listOfCurrencies...); err != nil {
			return nil, fmt.Errorf("failed to add currency to collection: %w", err)
		}
	}

	return cc, nil
}

// MustNewCurrencyCollection takes one or more list of currencies and returns them as a collection.
//
// With currencyStandard you can limit the collection to only allow currencies of a specific standard.
// Set it to "" if you want to allow multiple currency standards.
//
// It will panic on any error.
func MustNewCurrencyCollection(name string, currencyStandard string, listsOfCurrencies ...[]Currency) *CurrencyCollection {
	cc, err := NewCurrencyCollection(name, currencyStandard, listsOfCurrencies...)
	if err != nil {
		panic(fmt.Sprintf("failed to create currency collection %q: %v", name, err))
	}

	return cc
}

// Name returns the name of the currency collection.
func (cc *CurrencyCollection) Name() string {
	return cc.name
}

// Name returns the name of the currency standard.
// If this collection can contain multiple standards, this will return "".
func (cc *CurrencyCollection) CurrencyStandard() string {
	return cc.currencyStandard
}

// All returns the full list of currencies that are contained in this collection.
func (cc *CurrencyCollection) All() []Currency {
	cc.mutex.RLock()
	defer cc.mutex.RUnlock()

	return cc.all
}

// ByUniqueID finds a currency by its unique ID (e.g. 42170978).
func (cc *CurrencyCollection) ByUniqueID(uniqueID int32) Currency {
	cc.mutex.RLock()
	defer cc.mutex.RUnlock()

	return cc.byUniqueID[uniqueID]
}

// ByUniqueCode finds a currency by its unique code (e.g. "ISO4217-EUR").
func (cc *CurrencyCollection) ByUniqueCode(uniqueCode string) Currency {
	cc.mutex.RLock()
	defer cc.mutex.RUnlock()

	return cc.byUniqueCode[uniqueCode]
}

// ByNumericCode finds a currency by its numeric code (e.g. 978).
// This may not yield a result, as the numeric code is not unique across currency standards.
// Best is to use it only in combination with a collection of a single standard, like ISO4217Currencies.
func (cc *CurrencyCollection) ByNumericCode(numericCode int) Currency {
	if cc.byNumericCode != nil {
		cc.mutex.RLock()
		defer cc.mutex.RUnlock()

		return cc.byNumericCode[numericCode]
	}
	return nil
}

// ByCode finds a currency by its code (e.g. "EUR").
// This may not yield a result, as the code is not unique across currency standards.
// Best is to use it only in combination with a collection of a single standard, like ISO4217Currencies.
func (cc *CurrencyCollection) ByCode(code string) Currency {
	if cc.byCode != nil {
		cc.mutex.RLock()
		defer cc.mutex.RUnlock()

		return cc.byCode[code]
	}
	return nil
}

// add adds a currency to this collection.
func (cc *CurrencyCollection) add(c Currency) error {
	currencyStandard, uniqueID, uniqueCode, numericCode, code := c.Standard(), c.UniqueID(), c.UniqueCode(), c.NumericCode(), c.Code()

	// Check currency is allowed to be added.
	if cc.currencyStandard != "" && cc.currencyStandard != currencyStandard {
		return fmt.Errorf("currency standard of %q not allowed in this collection, only %q", currencyStandard, cc.currencyStandard)
	}

	// Check if the currency already exists.
	// Ignore duplicate entries, but prevent collisions of unique IDs, unique codes or codes.
	currencyByUniqueID, foundByUniqueID := cc.byUniqueID[uniqueID]
	currencyByUniqueCode, foundByUniqueCode := cc.byUniqueCode[uniqueCode]
	var currencyByCode Currency
	var foundByCode bool
	if cc.byCode != nil {
		currencyByCode, foundByCode = cc.byCode[code]
	}
	var currencyByNumericCode Currency
	var foundByNumericCode bool
	if cc.byNumericCode != nil {
		currencyByNumericCode, foundByNumericCode = cc.byNumericCode[numericCode]
	}

	if foundByUniqueID && currencyByUniqueID != c {
		// There is another currency with the same unique ID.
		return fmt.Errorf("currency %s has the same unique ID %d as the already existing currency %s", helperCurrencyUniqueCode(c), uniqueID, helperCurrencyUniqueCode(currencyByUniqueID))
	} else if foundByUniqueCode && currencyByUniqueCode != c {
		// There is another currency with the same unique code.
		return fmt.Errorf("currency with unique ID %d has the same unique code %q as the already existing currency with unique ID %d", uniqueID, uniqueCode, currencyByUniqueCode.UniqueID())
	} else if foundByNumericCode && currencyByNumericCode != c {
		// There is another currency with the same numeric code.
		return fmt.Errorf("currency %s has the same numeric code %d as the already existing currency %s", helperCurrencyUniqueCode(c), numericCode, helperCurrencyUniqueCode(currencyByNumericCode))
	} else if foundByCode && currencyByCode != c {
		// There is another currency with the same code.
		return fmt.Errorf("currency with unique ID %d has the same code %q as the already existing currency with unique ID %d", uniqueID, code, currencyByCode.UniqueID())
	}

	if _, found := cc.hasCurrency[c]; found {
		// Currency is already in the set. Ignore it.
		return nil
	}

	cc.all = append(cc.all, c)
	cc.hasCurrency[c] = struct{}{}
	cc.byUniqueID[uniqueID] = c
	cc.byUniqueCode[uniqueCode] = c
	if cc.byNumericCode != nil {
		cc.byNumericCode[numericCode] = c
	}
	if cc.byCode != nil {
		cc.byCode[code] = c
	}

	return nil
}

// Add adds one or more currencies to this collection.
func (cc *CurrencyCollection) Add(currencies ...Currency) error {
	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	for _, c := range currencies {
		if err := cc.add(c); err != nil {
			return err
		}
	}

	return nil
}
