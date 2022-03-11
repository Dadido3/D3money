// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"
)

func TestJSONMarshalling(t *testing.T) {
	values := []Value{
		MustFromString("0"),
		MustFromString("0 ISO4217-EUR"),
		MustFromString("-12345.6789"),
		MustFromString("-12345.6789 ISO4217-EUR"),
		MustFromString("12345.6789"),
		MustFromString("12345.6789 ISO4217-USD"),
		MustFromString("3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 ISO4217-XXX"),
	}

	expectedMarshalledValues := [][]byte{
		[]byte(`{"Amount":"0"}`),
		[]byte(`{"Amount":"0","Currency":"ISO4217-EUR"}`),
		[]byte(`{"Amount":"-12345.6789"}`),
		[]byte(`{"Amount":"-12345.6789","Currency":"ISO4217-EUR"}`),
		[]byte(`{"Amount":"12345.6789"}`),
		[]byte(`{"Amount":"12345.6789","Currency":"ISO4217-USD"}`),
		[]byte(`{"Amount":"3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679","Currency":"ISO4217-XXX"}`),
	}

	// Marshall values.
	marshalledValues := make([][]byte, len(values))
	for i, value := range values {
		var err error
		if marshalledValues[i], err = json.Marshal(value); err != nil {
			t.Errorf("json.Marshal(%#v) failed: %v", value, err)
		}
	}

	// Check marshalled values.
	if len(expectedMarshalledValues) != len(marshalledValues) {
		t.Fatalf("Amount of expected values %d and marshalled values %d is not equal.", len(expectedMarshalledValues), len(marshalledValues))
	}
	for i, expected := range expectedMarshalledValues {
		if string(expected) != string(marshalledValues[i]) {
			t.Errorf("Expected marshalled value %q, got %q", string(expected), string(marshalledValues[i]))
		}
	}

	// Unmarshal values.
	unmarshalledValues := make([]Value, len(marshalledValues))
	for i, value := range marshalledValues {
		if err := json.Unmarshal(value, &unmarshalledValues[i]); err != nil {
			t.Errorf("json.Marshal(%q) failed: %v", value, err)
		}
	}

	// Check roundtrip.
	if len(values) != len(unmarshalledValues) {
		t.Fatalf("Amount of values %d and unmarshalled values %d is not equal", len(values), len(unmarshalledValues))
	}
	for i, value := range values {
		if equal, err := value.Equal(unmarshalledValues[i]); err != nil {
			t.Errorf("JSON roundtrip failed: %v", err)
		} else if !equal {
			t.Errorf("JSON roundtrip failed. Values %q and %q are not equal", value, unmarshalledValues[i])
		}
	}
}

func TestBinaryMarshalling(t *testing.T) {
	values := []Value{
		MustFromString("0"),
		MustFromString("0 ISO4217-EUR"),
		MustFromString("-12345.6789"),
		MustFromString("-12345.6789 ISO4217-EUR"),
		MustFromString("12345.6789"),
		MustFromString("12345.6789 ISO4217-USD"),
		MustFromString("3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 ISO4217-XXX"),
	}

	// Marshall values.
	marshalledValues := make([][]byte, len(values))
	for i, value := range values {
		var err error
		if marshalledValues[i], err = value.MarshalBinary(); err != nil {
			t.Errorf("%#v.MarshalBinary() failed: %v", value, err)
		}
	}

	// Unmarshal values.
	unmarshalledValues := make([]Value, len(marshalledValues))
	for i, value := range marshalledValues {
		if err := unmarshalledValues[i].UnmarshalBinary(value); err != nil {
			t.Errorf("unmarshalledValues[i].UnmarshalBinary(%q) failed: %v", value, err)
		}
	}

	// Check roundtrip.
	if len(values) != len(unmarshalledValues) {
		t.Fatalf("Amount of values %d and unmarshalled values %d is not equal", len(values), len(unmarshalledValues))
	}
	for i, value := range values {
		if equal, err := value.Equal(unmarshalledValues[i]); err != nil {
			t.Errorf("Binary roundtrip failed: %v", err)
		} else if !equal {
			t.Errorf("Binary roundtrip failed. Values %q and %q are not equal", value, unmarshalledValues[i])
		}
	}
}

func TestTextMarshalling(t *testing.T) {
	values := []Value{
		MustFromString("0"),
		MustFromString("0 ISO4217-EUR"),
		MustFromString("-12345.6789"),
		MustFromString("-12345.6789 ISO4217-EUR"),
		MustFromString("12345.6789"),
		MustFromString("12345.6789 ISO4217-USD"),
		MustFromString("3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 ISO4217-XXX"),
	}

	expectedMarshalledValues := [][]byte{
		[]byte(`0`),
		[]byte(`0 ISO4217-EUR`),
		[]byte(`-12345.6789`),
		[]byte(`-12345.6789 ISO4217-EUR`),
		[]byte(`12345.6789`),
		[]byte(`12345.6789 ISO4217-USD`),
		[]byte(`3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 ISO4217-XXX`),
	}

	// Marshall values.
	marshalledValues := make([][]byte, len(values))
	for i, value := range values {
		var err error
		if marshalledValues[i], err = value.MarshalText(); err != nil {
			t.Errorf("%#v.MarshalText() failed: %v", value, err)
		}
	}

	// Check marshalled values.
	if len(expectedMarshalledValues) != len(marshalledValues) {
		t.Fatalf("Amount of expected values %d and marshalled values %d is not equal.", len(expectedMarshalledValues), len(marshalledValues))
	}
	for i, expected := range expectedMarshalledValues {
		if string(expected) != string(marshalledValues[i]) {
			t.Errorf("Expected marshalled value %q, got %q", string(expected), string(marshalledValues[i]))
		}
	}

	// Unmarshal values.
	unmarshalledValues := make([]Value, len(marshalledValues))
	for i, value := range marshalledValues {
		if err := unmarshalledValues[i].UnmarshalText(value); err != nil {
			t.Errorf("unmarshalledValues[i].UnmarshalText(%q) failed: %v", value, err)
		}
	}

	// Check roundtrip.
	if len(values) != len(unmarshalledValues) {
		t.Fatalf("Amount of values %d and unmarshalled values %d is not equal", len(values), len(unmarshalledValues))
	}
	for i, value := range values {
		if equal, err := value.Equal(unmarshalledValues[i]); err != nil {
			t.Errorf("Text roundtrip failed: %v", err)
		} else if !equal {
			t.Errorf("Text roundtrip failed. Values %q and %q are not equal", value, unmarshalledValues[i])
		}
	}
}

func TestGobEncoding(t *testing.T) {
	values := []Value{
		MustFromString("0"),
		MustFromString("0 ISO4217-EUR"),
		MustFromString("-12345.6789"),
		MustFromString("-12345.6789 ISO4217-EUR"),
		MustFromString("12345.6789"),
		MustFromString("12345.6789 ISO4217-USD"),
		MustFromString("3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 ISO4217-XXX"),
	}

	// Encode values.
	encodedValues := make([]bytes.Buffer, len(values))
	for i, value := range values {
		enc := gob.NewEncoder(&encodedValues[i])
		var err error
		if enc.Encode(value); err != nil {
			t.Errorf("enc.Encode(%#v) failed: %v", value, err)
		}
	}

	// Decode values.
	decodedValues := make([]Value, len(encodedValues))
	for i, value := range encodedValues {
		dec := gob.NewDecoder(&value)
		if err := dec.Decode(&decodedValues[i]); err != nil {
			t.Errorf("dec.Decode() failed: %v", err)
		}
	}

	// Check roundtrip.
	if len(values) != len(decodedValues) {
		t.Fatalf("Amount of values %d and decoded values %d is not equal", len(values), len(decodedValues))
	}
	for i, value := range values {
		if equal, err := value.Equal(decodedValues[i]); err != nil {
			t.Errorf("Binary roundtrip failed: %v", err)
		} else if !equal {
			t.Errorf("Binary roundtrip failed. Values %q and %q are not equal", value, decodedValues[i])
		}
	}
}
