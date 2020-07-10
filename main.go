package main

func main() {
	for countryCode := range countryList {
		getToken(countryCode)
		getCountryTrends(countryCode)
	}
}
