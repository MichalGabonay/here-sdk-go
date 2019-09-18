package heremaps

import (
	"encoding/json"
	"fmt"
	"time"
)

// GeocodeOutput - represent output data structure
type GeocodeOutput struct {
	Response GeocodeResponse
}

type GeocodeResponse struct {
	// MetaInfo GeocodeResponseMetaInfo
	View []GeocodeResponseView
}

type GeocodeResponseMetaInfo struct {
	Timestamp time.Time
}

type GeocodeResponseView struct {
	Type   string `json:"_type"`
	ViewId int32
	Result []GeocodeResponseResult
}

type GeocodeResponseResult struct {
	Relevance    float32
	MatchLevel   string
	MatchQuality GeocodeResponseResultMatchQuality
	MatchType    string
	Location     GeocodeResponseResultLocation
}

type GeocodeResponseResultMatchQuality struct {
	City        float32
	Street      []float32
	HouseNumber float32
}

type GeocodeResponseResultLocation struct {
	LocationId      string
	LocationType    string
	DisplayPosition GeocodeResponseResultLocationCords
	// NavigationPosition GeocodeResponseResultLocationCords
	// MapView            GeocodeResponseResultLocationMapView
	Address GeocodeResponseResultLocationAddress
}

type GeocodeResponseResultLocationMapView struct {
	TopLeft     GeocodeResponseResultLocationCords
	BottomRight GeocodeResponseResultLocationCords
}

type GeocodeResponseResultLocationCords struct {
	Latitude  float32
	Longitude float32
}

type GeocodeResponseResultLocationAddress struct {
	Label          string
	Country        string
	State          string
	County         string
	City           string
	District       string
	Street         string
	HouseNumber    string
	PostalCode     string
	AdditionalData []GeocodeResponseResultLocationAddressAdditionalData
}

type GeocodeResponseResultLocationAddressAdditionalData struct {
	Value string `json:"Value"`
	Key   string `json:"Key"`
}

// Location - represent siple structure for location (geocode)
type Location struct {
	Addr  Address
	Cords Cords
}

// buildURLForGeocodeRequest - create url for GET request on Here maps API to search for location
func buildURLForGeocodeRequest(appID, appCode, input string) string {

	const BaseAPIURL = "https://geocoder.api.here.com/6.2/geocode.json"

	url := BaseAPIURL + "?app_id=" + appID + "&app_code=" + appCode + "&searchtext=" + input

	return url

}

// ParseGeocode - take array of bytes from API request and returne structured geocode
func parseGeocode(body []byte) (*GeocodeOutput, error) {
	var g = new(GeocodeOutput)
	err := json.Unmarshal(body, &g)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return g, err
}

// SimplifyLocation - make location data structutre sipler
func SimplifyLocation(geocode *GeocodeOutput) *Location {

	location := geocode.Response.View[0].Result[0].Location

	if len(geocode.Response.View) == 0 {
		return nil
	}
	if len(geocode.Response.View[0].Result) == 0 {
		return nil
	}

	addr := Address{
		Country:     location.Address.Country,
		State:       location.Address.State,
		County:      location.Address.County,
		City:        location.Address.City,
		District:    location.Address.District,
		Street:      location.Address.Street,
		HouseNumber: location.Address.HouseNumber,
		PostalCode:  location.Address.PostalCode,
	}

	cords := Cords{
		Lat: cordToString(location.DisplayPosition.Latitude),
		Lng: cordToString(location.DisplayPosition.Latitude),
	}

	loc := Location{addr, cords}

	return &loc
}

func cordToString(cord float32) string {
	s := fmt.Sprintf("%f", cord)
	return s
}
