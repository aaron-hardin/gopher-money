package format

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	actual, err := Format("USD", 44.3)
	if err != nil {
		t.Fatal(err)
	}
	expected := "$44.30"
	if actual != expected {
		t.Fatal(fmt.Sprintf("USD should be formatted as $<value> with 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual, err = Format("USD", -44.3)
	if err != nil {
		t.Fatal(err)
	}
	expected = "$-44.30"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-USD should be formatted as $-<value> with 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual, err = Format("JPY", 44.3)
	if err != nil {
		t.Fatal(err)
	}
	expected = "¥44"
	if actual != expected {
		t.Fatal(fmt.Sprintf("JPY should be formatted as ¥<value> with 0 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual, err = Format("JPY", -44.3)
	if err != nil {
		t.Fatal(err)
	}
	expected = "¥-44"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-JPY should be formatted as ¥-<value> with 0 decimal places, actual: %s, expected: %s", actual, expected))
	}
}

func TestFormatAs(t *testing.T) {
	actual, err := FormatAs("USD", 44.3, "%s %v")
	if err != nil {
		t.Fatal(err)
	}
	expected := "$ 44.30"
	if actual != expected {
		t.Fatal(fmt.Sprintf("USD should be formatted using format given $ and 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual, err = FormatAs("USD", -44.3, "%s (%v)")
	if err != nil {
		t.Fatal(err)
	}
	expected = "$ (44.30)"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-USD should be formatted using format given $ and 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual, err = FormatAs("JPY", 44.3, "%s %v")
	if err != nil {
		t.Fatal(err)
	}
	expected = "¥ 44"
	if actual != expected {
		t.Fatal(fmt.Sprintf("JPY should be formatted using format given ¥ and 0 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual, err = FormatAs("JPY", -44.3, "%s (%v)")
	if err != nil {
		t.Fatal(err)
	}
	expected = "¥ (44)"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-JPY should be formatted using format given ¥ and 0 decimal places, actual: %s, expected: %s", actual, expected))
	}
}
