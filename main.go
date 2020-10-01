package main

import (
	"fmt"

	"github.com/aaron-hardin/gopher-money/format"
	"github.com/aaron-hardin/gopher-money/money"
	"github.com/aaron-hardin/gopher-money/rates"
)

func main() {
	client := rates.NewApiClient("TODO: put key here")
	rates := client.GetRates()
	converter := money.NewConverter(rates)
	rawValue := converter.Convert("USD", "JPY", 44.3)
	fmt.Println(rawValue)
	fmt.Println(format.Round("JPY", rawValue))
	fmt.Println(format.Format("JPY", rawValue))
	fmt.Println(format.FormatAs("JPY", -rawValue, "%s(%v)"))
}
