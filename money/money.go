package money

func (c *Converter) Convert(from string, to string, value float64) float64 {
	return ConvertWithRates(c.Rates, from, to, value)
}

func ConvertWithRates(rates map[string]float64, from string, to string, value float64) float64 {
	return value * getRate(rates, from, to)
}

func getRate(rates map[string]float64, from string, to string) float64 {
	// TODO: need to have error handling here
	return rates[to] / rates[from]
}

func NewConverter(rates map[string]float64) *Converter {
	c := new(Converter)
	c.Rates = rates
	return c
}

type Converter struct {
	Rates map[string]float64
}
