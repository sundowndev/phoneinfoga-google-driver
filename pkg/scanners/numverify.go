package scanners

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/parnurzeal/gorequest"
	"github.com/sundowndev/phoneinfoga/pkg/utils"
	// "golang.org/x/net/html"
)

// Numverify REST API response
type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

// NumverifyScan fetches Numverify's API
func NumverifyScan(number string) (res *Numverify, err error) {
	utils.LoggerService.Infoln("Running Numverify.com scan...")

	_, _, errs := gorequest.New().Get("http://numverify.com/").End()
	if errs != nil {
		log.Fatal(errs)
	}

	// doc, err := html.Parse(body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(doc)

	// Then fetch infos
	apiKey := md5.New()
	io.WriteString(apiKey, "1")
	safeNumber := url.QueryEscape(number)

	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=%s&number=%s", apiKey.Sum(nil), safeNumber)

	// Build the request
	resp2, err := http.Get(url)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// Fill the response with the data from the JSON
	var response Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp2.Body).Decode(&response); err != nil {
		log.Println(err)
	}

	res = &Numverify{
		Valid:               response.Valid,
		Number:              response.Number,
		LocalFormat:         response.LocalFormat,
		InternationalFormat: response.InternationalFormat,
		CountryPrefix:       response.CountryPrefix,
		CountryCode:         response.CountryCode,
		CountryName:         response.CountryName,
		Location:            response.Location,
		Carrier:             response.Carrier,
		LineType:            response.LineType,
	}

	return res, nil
}
