package country

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// Country struct
type Country struct {
	data map[string]interface{}
}

// NewCountry will create a new instance of country struct
func NewCountry() *Country {
	instance := &Country{}

	instance = instance.readCountriesDataFile()

	return instance
}

// Read countries data file
func (country *Country) readCountriesDataFile() *Country {
	if len(country.data) > 0 {
		return country
	}

	dataPath, _ := filepath.Abs("./data/countries.json")

	file, _ := ioutil.ReadFile(dataPath)

	json.Unmarshal([]byte(file), &country.data)

	return country
}

// All will return all countries name and dialing code
func (country *Country) All() map[string]interface{} {
	return country.readCountriesDataFile().data
}

// Get a single country by the country ISO 3166-1 Alpha-2 code
func (country *Country) getCountry(code string) (interface{}, error) {
	details := country.readCountriesDataFile().data[code]

	if details == nil {
		return nil, NewValidationError(code)
	}

	return details, nil
}

// Get a country name and dialing code by the country ISO 3166-1 Alpha-2 code
func (country *Country) Get(code interface{}) (interface{}, error) {
	switch code.(type) {
	case string:
		return country.getCountry(code.(string))
	case []string:
		data, err := country.gets(code.([]string))

		return data, err
	default:
		return nil, nil
	}
}

// GetName will return a country name by the country ISO 3166-1 Alpha-2 code
func (country *Country) GetName(code string) (interface{}, error) {
	data, err := country.getCountry(code)

	if err != nil {
		return nil, err
	}

	details, _ := data.(map[string]interface{})

	return details["name"], nil
}

// GetDialingCode will return a country dialing code by the country ISO 3166-1 Alpha-2 code
func (country *Country) GetDialingCode(code string) (interface{}, error) {
	data, err := country.getCountry(code)

	if err != nil {
		return nil, err
	}

	details, _ := data.(map[string]interface{})

	return details["code"], nil
}

// Get countries name and dialing code by the country ISO 3166-1 Alpha-2 codes
func (country *Country) gets(codes []string) (map[string]interface{}, error) {
	countries := make(map[string]interface{})

	for _, code := range codes {
		details, err := country.getCountry(code)

		if err != nil {
			return nil, err
		}

		countries[code] = details
	}

	return countries, nil
}
