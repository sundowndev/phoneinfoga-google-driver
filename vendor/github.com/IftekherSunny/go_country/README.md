### Country

[![Project status](https://img.shields.io/badge/release-v1.0-green.svg)](https://github.com/IftekherSunny/go_country/releases)
[![Build Status](https://travis-ci.org/IftekherSunny/go_country.svg?branch=master)](https://travis-ci.org/IftekherSunny/go_country)
[![cover.run go](https://cover.run/go/github.com/IftekherSunny/go_country.svg)](https://cover.run/go/github.com/IftekherSunny/go_country)
[![Go Report Card](https://goreportcard.com/badge/github.com/IftekherSunny/go_country)](https://goreportcard.com/report/github.com/IftekherSunny/go_country)
[![GoDoc](https://godoc.org/github.com/IftekherSunny/go_country?status.svg)](https://godoc.org/github.com/IftekherSunny/go_country)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/IftekherSunny/go_country/blob/master/LICENSE)

Country is the package that helps you to get country name and dialing code by the country ISO 3166-1 Alpha-2 code.

### Installation Process

```
 go get github.com/IftekherSunny/go_country
```

### Basic Uses

##### Get all countries name and dialing code

```go
country := country.NewCountry()

countries := country.All()
```

##### Get a country name and dialing code

```go
country := country.NewCountry()

countryDetails, _ := country.Get("BD")
```

##### Get multiple countries name and dialing code

```go
country := country.NewCountry()

countries, _ := country.Get([]string{"BD", "US"})
```

##### Get a country name

```go
country := country.NewCountry()

name, _ := country.GetName("BD")
```

##### Get a country dialing code

```go
country := country.NewCountry()

dialingCode, _ := country.GetDialingCode("BD")
```

### Test

##### Run tests

```
go test -v
```

### License
This package is licensed under the [MIT License](https://github.com/iftekhersunny/country-for-golang/blob/master/LICENSE)
