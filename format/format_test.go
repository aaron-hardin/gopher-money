package format

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	actual := Format("USD", 44.3)
	expected := "$44.30"
	if actual != expected {
		t.Fatal(fmt.Sprintf("USD should be formatted as $<value> with 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual = Format("USD", -44.3)
	expected = "$-44.30"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-USD should be formatted as $-<value> with 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual = Format("JPY", 44.3)
	expected = "¥44"
	if actual != expected {
		t.Fatal(fmt.Sprintf("JPY should be formatted as ¥<value> with 0 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual = Format("JPY", -44.3)
	expected = "¥-44"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-JPY should be formatted as ¥-<value> with 0 decimal places, actual: %s, expected: %s", actual, expected))
	}
}

func TestFormatAs(t *testing.T) {
	actual := FormatAs("USD", 44.3, "%s %v")
	expected := "$ 44.30"
	if actual != expected {
		t.Fatal(fmt.Sprintf("USD should be formatted using format given $ and 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual = FormatAs("USD", -44.3, "%s (%v)")
	expected = "$ (44.30)"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-USD should be formatted using format given $ and 2 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual = FormatAs("JPY", 44.3, "%s %v")
	expected = "¥ 44"
	if actual != expected {
		t.Fatal(fmt.Sprintf("JPY should be formatted using format given ¥ and 0 decimal places, actual: %s, expected: %s", actual, expected))
	}

	actual = FormatAs("JPY", -44.3, "%s (%v)")
	expected = "¥ (44)"
	if actual != expected {
		t.Fatal(fmt.Sprintf("-JPY should be formatted using format given ¥ and 0 decimal places, actual: %s, expected: %s", actual, expected))
	}
}
