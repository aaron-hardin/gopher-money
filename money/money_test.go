package money

import (
	"testing"
)

func TestConversion(t *testing.T) {
	rates := make(map[string]float64)
	rates["USD"] = 1 // This is the 'base'
	rates["JPY"] = 100
	convertedValue := ConvertWithRates(rates, "USD", "USD", 43.23)
	if convertedValue != 43.23 {
		t.Fatal("No conversion should result in same number")
	}
	convertedValue = ConvertWithRates(rates, "USD", "JPY", 44.3)
	if convertedValue != 4430 {
		t.Fatal("Conversion should result in different number")
	}
}
