// gopher-money is a tool for all things money.
//
// Usage:
//
// 	gopher-money <command> [arguments]
//
// The commands are:
//
// 	convert     convert from one currency to another
// 	round       round currency to nearest whole amount
// 	format      format a currency amount
//
// Use "gopher-money <command> --help" for more information about a command.
//
//
// Convert from one currency to another
//
// Usage:
//
// 	gopher-money convert <from currency code> <to currency code> <value to convert> [--apiKey=<apiKey>]
//
// Convert uses current exchange rates to convert from one currency to another.
//
// The apiKey is optional and defaults to the environment variable GOPHER_MONEY_API_KEY
//
//
// Round currency to nearest whole amount
//
// Usage:
//
// 	gopher-money round <currency> <value>
//
// Round takes an input value and rounds to the nearest whole unit according to the currency code.
//
// Examples:
//
//  USD
//    The standard unit for US dollar is 1 cent so it will be rounded to nearest .01
//  JPY
//    The standard unit for Japanese Yen is 1 yen so it will be rounded to the nearest 1
//
//
// Format a currency amount
//
// Usage:
//
// 	gopher-money format <currency> <value> [format]
//
// Format takes an input value and formats it according to the currency code.
// The value is first rounded in the way that the round command is applied.
// The result is then formatted according to the format string with a default of %s%v.
// %v specifies the value after rounding
// %s specifies the currency symbol
//
package main
