package heremaps

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HereClient defines HERE map client
type HereClient struct {
	AppID   string
	AppCode string
}

// Address represent address returned from HERE maps API
type Address struct {
	Country     string
	State       string
	County      string
	City        string
	District    string
	Street      string
	HouseNumber string
	PostalCode  string
}

// Cords represent coordinates of location
type Cords struct {
	Lat string
	Lng string
}

type RouteSummary struct {
	Distance uint32 // in meters
	Duration uint32 // in seconds
}

// NewClient returns a new Medium API client which can be used to make RPC requests.
func NewClient(id, code string) *HereClient {
	return &HereClient{
		AppID:   id,
		AppCode: code,
	}
}

// SearchLocation - returns structured address and coordinates for location from address written as string
func (hc *HereClient) SearchLocation(input string) (*GeocodeOutput, error) {
	url := buildURLForGeocodeRequest(hc.AppID, hc.AppCode, input)

	log.Print("url: " + url)

	response, err := http.Get(url)
	if err != nil {
		log.Print("err1")
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print("err2")
		return nil, err
	}

	geocode, err := parseGeocode(body)
	if err != nil {
		log.Print("err3")
		return nil, err
	}

	return geocode, err
}

func (hc *HereClient) CalculateRoute(routingInput *RoutingInput) (*RoutingOutput, error) {
	url := buildURLForRoutingRequest(hc.AppID, hc.AppCode, routingInput)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	route, err := parseRoute(body)
	if err != nil {
		return nil, err
	}

	return route, err

}
