package main

import (
	"fmt"
	"log"
	"soap/myservice"

	"github.com/hooklift/gowsdl/soap"
)

func main() {

	countryServiceClient := soap.NewClient("http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso")
	countryService := myservice.NewCountryInfoServiceSoapType(countryServiceClient)

	fullCountryInfoRequest := &myservice.CountriesUsingCurrency{SISOCurrencyCode: "USD"}
	fullCountryInfoResponse, err := countryService.CountriesUsingCurrency(fullCountryInfoRequest)
	if err != nil {
		log.Fatalf("Error calling FullCountryInfo: %v", err)
	}

	for i := range fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName {
		fmt.Printf("Country Code: %s, Country Name: %s\n",
			fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SISOCode,
			fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SName)
	}

}
