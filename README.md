![Test](https://github.com/aaron-hardin/gopher-money/workflows/Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/aaron-hardin/gopher-money?status.svg)](https://godoc.org/github.com/aaron-hardin/gopher-money)

# gopher-money
Command line utility for all things money
```shell script
gopher-money convert USD JPY 44.3
> 4679.36746875
gopher-money round JPY 44.3
> 44
gopher-money format USD 44.3
> $44.30
```

Install using the following command
```
go get github.com/aaron-hardin/gopher-money
```

# gopher-money - rates [![GoDoc](https://godoc.org/github.com/aaron-hardin/gopher-money/rates?status.svg)](https://godoc.org/github.com/aaron-hardin/gopher-money/rates)
Wrapper around Open Exchange Rates to provide a convenient way to get current exchange rates
```go
client := rates.NewApiClient(apiKey)
// rates is map[string]float64
// mapping currency code to current exchange rate
rates, err := client.GetRates()
```

# gopher-money - money [![GoDoc](https://godoc.org/github.com/aaron-hardin/gopher-money/money?status.svg)](https://godoc.org/github.com/aaron-hardin/gopher-money/money)
Provides functionality for using rates to convert between currencies
```go
converter := money.NewConverter(rates)
rawValue, err := converter.Convert(fromCurrencyCode, toCurrencyCode, value)
```

# gopher-money - format [![GoDoc](https://godoc.org/github.com/aaron-hardin/gopher-money/format?status.svg)](https://godoc.org/github.com/aaron-hardin/gopher-money/format)
Provides functionality for formatting currencies, allows converting to correct precision as well as formatting to string with symbol

When providing a custom format string, %v is used for the value and %s is used for the currency symbol
```go
roundedValue, err := format.Round(currencyCode, value)
displayValue, err := format.Format(currencyCode, value)
displayValue, err = format.FormatAs(currencyCode, value, customFormat)
```

# TODO
* Add more features to rates library to support more features of Open Exchange Rates
