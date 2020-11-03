package main

import (
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"log"
	"os"

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
	Cache  bool    `help:"Uses local cache for exchange rates."`
}

const cacheFile = "gopher-money-rates-cache.txt"

func (c *ConvertCmd) Run(ctx *Context) error {
	var exchangeRates map[string]float64
	saveRates := c.Cache
	callApi := true
	if c.Cache {
		if _, err := os.Stat(cacheFile); err == nil {
			// cache file exists
			fi, err := os.Open(cacheFile)
			if err != nil {
				return err
			}
			defer fi.Close()

			fz, err := gzip.NewReader(fi)
			if err != nil {
				return err
			}
			defer fz.Close()

			decoder := gob.NewDecoder(fz)

			// Decoding the serialized data
			err = decoder.Decode(&exchangeRates)
			if err != nil {
				return err
			}

			// since we loaded from file we don't need to call api
			callApi = false

			// since we loaded from file we don't need to save to the file
			saveRates = false
		} else if os.IsNotExist(err) {
			callApi = true
		} else {
			// could not read the file, but not because it wasn't there
			return err
		}
	} else {
		callApi = true
	}

	if callApi {
		client := rates.NewApiClient(c.ApiKey)
		var err error // Declaring here since we can't use := syntax for exchangeRates
		exchangeRates, err = client.GetRates()
		if err != nil {
			return err
		}
	}

	if saveRates {
		fi, err := os.Create(cacheFile)
		if err != nil {
			return err
		}
		defer fi.Close()

		fz := gzip.NewWriter(fi)
		defer fz.Close()

		encoder := gob.NewEncoder(fz)
		err = encoder.Encode(exchangeRates)
		if err != nil {
			return err
		}
	}

	converter := money.NewConverter(exchangeRates)
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
