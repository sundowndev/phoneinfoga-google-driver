package utils

import (
	"os"
	"regexp"

	phoneiso3166 "github.com/onlinecity/go-phone-iso3166"
)

// FormatNumber formats a phone number to remove
// unnecessary chars and avoid dealing with unwanted input.
func FormatNumber(n string) string {
	re := regexp.MustCompile(`[_\W]+`)
	number := re.ReplaceAllString(n, "")

	re = regexp.MustCompile("^[0-9]+$")

	if len(re.FindString(number)) == 0 {
		LoggerService.Errorln("Number is not valid.")
		os.Exit(0)
	}

	return number
}

// ParseCountryCode parses a phone number and returns ISO country code.
// This is required in order to use the phonenumbers library.
func ParseCountryCode(n string) string {
	return phoneiso3166.E164.LookupString(FormatNumber(n))
}
