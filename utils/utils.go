package utils

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	phoneiso3166 "github.com/onlinecity/go-phone-iso3166"
	"github.com/sundowndev/go-covermyass/utils"
)

// FormatNumber ...
func FormatNumber(n string) (num string) {
	num = strings.Replace(n, "+", "", 1)

	re := regexp.MustCompile("^[0-9]+$")

	if len(re.FindString(num)) == 0 {
		utils.LoggerService.Error("Number is not valid.")
		os.Exit(0)
	}

	return num
}

// ParseCountryCode ...
func ParseCountryCode(n string) string {
	num, err := strconv.ParseUint(FormatNumber(n), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return phoneiso3166.E164.Lookup(num)
}
