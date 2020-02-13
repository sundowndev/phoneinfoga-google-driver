package scanners

import (
	"os"

	"github.com/nyaruka/phonenumbers"
	"github.com/sundowndev/phoneinfoga/utils"
)

// LocalScan performs a local scan of a phone
// number using phonenumbers library
func LocalScan(number string) *Number {
	country := utils.ParseCountryCode(number)

	num, err := phonenumbers.Parse(number, country)

	if err != nil {
		utils.LoggerService.Errorln("The number is not valid")
		os.Exit(0)
	}

	res := &Number{
		local:         phonenumbers.Format(num, phonenumbers.NATIONAL),
		international: phonenumbers.Format(num, phonenumbers.E164),
		countryCode:   num.GetCountryCode(),
		carrier:       num.GetPreferredDomesticCarrierCode(),
	}

	return res
}
