package scanners

import (
	"github.com/sundowndev/dorkgen"
	"github.com/sundowndev/phoneinfoga/pkg/utils"
)

// GoogleSearchDork is the common format for dork requests
type GoogleSearchDork struct {
	Number string `json:"number"`
	Dork   string `json:"dork"`
	URL    string
}

// GoogleSearchScan ...
func GoogleSearchScan(number *Number) []*GoogleSearchDork {
	utils.LoggerService.Infoln("Running Google search scan...")

	dorks := []*dorkgen.GoogleSearch{
		// TODO: Disposable phone numbers
		// TODO: Individuals
		(&dorkgen.GoogleSearch{}).
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		// Social medias
		(&dorkgen.GoogleSearch{}).
			Site("facebook.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("twitter.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("linkedin.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("instagram.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		// Reputation
		// (&dorkgen.GoogleSearch{}).
		// 	Site("whosenumber.info").
		// 	Intext(number.E164).
		// 	Intitle("who called"),
		(&dorkgen.GoogleSearch{}).
			// Intitle("Phone Fraud").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("findwhocallsme.com").
			Intext(number.E164).
			Or().
			Intext(number.International),
		(&dorkgen.GoogleSearch{}).
			Site("yellowpages.ca").
			Intext(number.E164),
		(&dorkgen.GoogleSearch{}).
			Site("phonenumbers.ie").
			Intext(number.E164),
		(&dorkgen.GoogleSearch{}).
			Site("who-calledme.com").
			Intext(number.E164),
		(&dorkgen.GoogleSearch{}).
			Site("usphonesearch.net").
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("whocalled.us").
			Inurl(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("quinumero.info").
			Intext(number.Local).
			Or().
			Intext(number.International),
		(&dorkgen.GoogleSearch{}).
			Site("uk.popularphotolook.com").
			Inurl(number.Local),
	}

	results := []*GoogleSearchDork{}

	for _, dork := range dorks {
		results = append(results, &GoogleSearchDork{
			Number: number.E164,
			Dork:   dork.ToString(),
			URL:    dork.ToURL(),
		})
	}

	return results
}
