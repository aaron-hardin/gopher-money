package format

import "testing"

func TestFormat(t *testing.T) {
	data := []struct {
		curr string
		in   float64
		out  string
	}{
		{"USD", 44.3, "$44.30"},
		{"USD", -44.3, "$-44.30"},
		{"JPY", 44.3, "¥44"},
		{"JPY", -44.3, "¥-44"},
	}

	for _, exp := range data {
		res, err := Format(exp.curr, exp.in)
		if err != nil {
			t.Fatal(err)
		}
		if res != exp.out {
			t.Fatalf("'%v' didn't match '%v'\n", res, exp.out)
		}
	}
}

func TestFormatAs(t *testing.T) {
	data := []struct {
		curr   string
		in     float64
		format string
		out    string
	}{
		{"USD", 44.3, "%s %v", "$ 44.30"},
		{"USD", -44.3, "%s (%v)", "$ (44.30)"},
		{"JPY", 44.3, "%s %v", "¥ 44"},
		{"JPY", -44.3, "%s (%v)", "¥ (44)"},
	}

	for _, exp := range data {
		res, err := FormatAs(exp.curr, exp.in, exp.format)
		if err != nil {
			t.Fatal(err)
		}
		if res != exp.out {
			t.Fatalf("'%v' didn't match '%v'\n", res, exp.out)
		}
	}
}
