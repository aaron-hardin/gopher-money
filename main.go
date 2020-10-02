package main

import (
	"fmt"
	"log"

	"github.com/aaron-hardin/gopher-money/format"
	"github.com/aaron-hardin/gopher-money/money"
	"github.com/aaron-hardin/gopher-money/rates"
)

func main() {
	client := rates.NewApiClient("TODO: put key here")
	rates, err := client.GetRates()
	if err != nil {
		log.Fatal(err)
	}

	converter := money.NewConverter(rates)
	rawValue, err := converter.Convert("USD", "JPY", 44.3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawValue)

	roundedValue, err := format.Round("JPY", rawValue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(roundedValue)

	displayValue, err := format.Format("JPY", rawValue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(displayValue)

	displayValue, err = format.FormatAs("JPY", -rawValue, "%s(%v)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(displayValue)
}
