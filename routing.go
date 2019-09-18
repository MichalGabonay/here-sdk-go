package heremaps

import (
	"encoding/json"
	"fmt"
)

// RoutingInput - represent input data structure
type RoutingInput struct {
	Origin      Cords
	Destination Cords
}

// RoutingOutput - represent output data structure
type RoutingOutput struct {
	Response RoutingResponse `json:"response"`
}

type RoutingResponse struct {
	Route []RoutingResponseRoute `json:"route"`
}

type RoutingResponseRoute struct {
	Summary RoutingResponseRouteSummary `json:"summary"`
}

type RoutingResponseRouteSummary struct {
	Distance uint32 `json:"distance"` // in meters
	BaseTime uint32 `json:"baseTime"` // in seconds
}

// buildURLForGeocodeRequest - create url for GET request on Here maps API to search for location
func buildURLForRoutingRequest(appID, appCode string, input *RoutingInput) string {

	const BaseAPIURL = "https://route.api.here.com/routing/7.2/calculateroute.json"

	mode := "&mode=balanced;truck"

	origin := "&waypoint0=geo!" + input.Origin.Lat + "," + input.Origin.Lng
	destination := "&waypoint1=geo!" + input.Destination.Lat + "," + input.Destination.Lng

	url := BaseAPIURL + "?app_id=" + appID + "&app_code=" + appCode + origin + destination + mode

	return url
}

// parseRoute - take array of bytes from API request and returne structured route info
func parseRoute(body []byte) (*RoutingOutput, error) {
	var g = new(RoutingOutput)
	err := json.Unmarshal(body, &g)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return g, err
}

// SummaryOfRoute - return summary of route - distance in meters and duration in seconds
func SummaryOfRoute(routing *RoutingOutput) *RouteSummary {

	if len(routing.Response.Route) == 0 {
		return nil
	}

	summary := RouteSummary{
		Distance: routing.Response.Route[0].Summary.Distance,
		Duration: routing.Response.Route[0].Summary.BaseTime,
	}

	return &summary
}
