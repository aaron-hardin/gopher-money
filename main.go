package main

import (
	"fmt"

	"github.com/aaron-hardin/gopher-money/money"
	"github.com/aaron-hardin/gopher-money/rates"
)

func main() {
	client := rates.NewApiClient("TODO: put key here")
	rates := client.GetRates()
	converter := money.NewConverter(rates)
	fmt.Println(converter.Convert("USD", "JPY", 44.3))
}
