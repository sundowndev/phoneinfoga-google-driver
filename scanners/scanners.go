package scanners

// Number is a phone number
type Number struct {
	local         string
	international string
	countryCode   int32
	carrier       string
}
