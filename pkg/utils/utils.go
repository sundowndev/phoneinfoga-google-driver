package utils

import (
	"os"
	"regexp"
	"strings"

	phoneiso3166 "github.com/onlinecity/go-phone-iso3166"
)

// FormatNumber formats a phone number to remove
// unnecessary chars and avoid dealing with unwanted input.
func FormatNumber(n string) string {
	n = strings.ReplaceAll(n, "+", "")
	n = strings.ReplaceAll(n, "-", "")
	n = strings.ReplaceAll(n, "(", "")
	n = strings.ReplaceAll(n, ")", "")
	n = strings.ReplaceAll(n, " ", "")

	re := regexp.MustCompile("^[0-9]+$")

	if len(re.FindString(n)) == 0 {
		LoggerService.Errorln("Number is not valid.")
		os.Exit(0)
	}

	return n
}

// ParseCountryCode parses a phone number and returns ISO country code.
// This is required in order to use the phonenumbers library.
func ParseCountryCode(n string) string {
	return phoneiso3166.E164.LookupString(FormatNumber(n))
}
