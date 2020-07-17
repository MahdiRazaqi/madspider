package main

import "github.com/MahdiRazaqi/madspider/madspider"

func main() {
	for countryCode := range madspider.CountryList {
		madspider.GetToken(countryCode)
		madspider.GetCountryTrends(countryCode)
	}
}
