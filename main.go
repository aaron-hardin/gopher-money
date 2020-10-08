package main

import (
	"fmt"
	"log"

	"github.com/aaron-hardin/gopher-money/format"
	"github.com/aaron-hardin/gopher-money/money"
	"github.com/aaron-hardin/gopher-money/rates"
	"github.com/alecthomas/kong"
)

type Context struct {
	Debug bool
}

type ConvertCmd struct {
	ApiKey string  `name:"apiKey" help:"API key for Open Exchange Rates." env:"GOPHER_MONEY_API_KEY"`
	From   string  `arg name:"from" help:"From currency."`
	To     string  `arg name:"to" help:"To currency."`
	Value  float64 `arg name:"value" help:"Value to convert."`
}

func (c *ConvertCmd) Run(ctx *Context) error {
	client := rates.NewApiClient(c.ApiKey)
	rates, err := client.GetRates()
	if err != nil {
		return err
	}

	converter := money.NewConverter(rates)
	rawValue, err := converter.Convert(c.From, c.To, c.Value)
	if err != nil {
		return err
	}

	fmt.Println(rawValue)

	return nil
}

type RoundCmd struct {
	Currency string  `arg name:"currency" help:"Currency code to format as."`
	Value    float64 `arg name:"value" help:"Value to format."`
}

func (r *RoundCmd) Run(ctx *Context) error {
	roundedValue, err := format.Round(r.Currency, r.Value)
	if err != nil {
		return err
	}
	fmt.Println(roundedValue)

	return nil
}

type FormatCmd struct {
	Currency string  `arg name:"currency" help:"Currency code to format as."`
	Value    float64 `arg name:"value" help:"Value to format."`
	Format   string  `arg optional name:"format" help:"Format to use %v for value and %s for currency symbol."`
}

func (f *FormatCmd) Run(ctx *Context) error {
	if f.Format == "" {
		displayValue, err := format.Format(f.Currency, f.Value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(displayValue)
	} else {
		displayValue, err := format.FormatAs(f.Currency, f.Value, f.Format)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(displayValue)
	}
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Convert ConvertCmd `cmd help:"Convert currency."`
	Round   RoundCmd   `cmd help:"Round currency to nearest whole amount."`
	Format  FormatCmd  `cmd help:"Format currency."`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
