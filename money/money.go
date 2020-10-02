package money

import "fmt"

func (c *Converter) Convert(from string, to string, value float64) (float64, error) {
	return ConvertWithRates(c.Rates, from, to, value)
}

func ConvertWithRates(rates map[string]float64, from string, to string, value float64) (float64, error) {
	rate, err := getRate(rates, from, to)
	if err != nil {
		return 0, err
	}

	return value * rate, nil
}

func getRate(rates map[string]float64, from string, to string) (float64, error) {
	if rateTo, ok := rates[to]; ok {
		if rateFrom, ok := rates[from]; ok {
			return rateTo / rateFrom, nil
		}
		return 0, fmt.Errorf("money: rates does not contain value for currency code: %v", from)
	}
	return 0, fmt.Errorf("money: rates does not contain value for currency code: %v", to)
}

func NewConverter(rates map[string]float64) *Converter {
	c := new(Converter)
	c.Rates = rates
	return c
}

type Converter struct {
	Rates map[string]float64
}
