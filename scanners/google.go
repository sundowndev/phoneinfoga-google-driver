package scanners

import (
	"github.com/sundowndev/dorkgen"
	"github.com/sundowndev/phoneinfoga/utils"
)

// GoogleSearchDork ...
type GoogleSearchDork struct {
	number string
	URL    string
}

// GoogleSearchScan ...
func GoogleSearchScan(number *Number) []string {
	utils.LoggerService.Infoln("Running Google search scan...")

	res := []string{
		(&dorkgen.GoogleSearch{}).
			Site("facebook.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local).
			ToURL(),
		(&dorkgen.GoogleSearch{}).
			Site("twitter.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local).
			ToURL(),
		(&dorkgen.GoogleSearch{}).
			Site("linkedin.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local).
			ToURL(),
		(&dorkgen.GoogleSearch{}).
			Site("instagram.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local).
			ToURL(),
	}

	return res
}
