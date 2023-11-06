package trackingmore

import (
	"context"
	"errors"
	"net/http"
	"regexp"
)

type AirWaybillItem struct {
	AwbNumber        string   `json:"awb_number"`
	StatusNumber     string   `json:"status_number"`
	Weight           string   `json:"weight"`
	Piece            string   `json:"piece"`
	Origin           string   `json:"origin"`
	Destination      string   `json:"destination"`
	FlightWayStation []string `json:"flight_way_station"`
	LastEvent        string   `json:"last_event"`
	FlightInfoNew    []struct {
		FlightNumber    string `json:"flight_number"`
		DepartStation   string `json:"depart_station"`
		ArrivalStation  string `json:"arrival_station"`
		PlanDepartTime  string `json:"plan_depart_time"`
		DepartTime      string `json:"depart_time"`
		PlanArrivalTime string `json:"plan_arrival_time"`
		ArrivalTime     string `json:"arrival_time"`
		Piece           string `json:"piece"`
		Weight          string `json:"weight"`
		Status          string `json:"status"`
	} `json:"flight_info_new"`
	FlightInfo map[string]FlightDetails `json:"flight_info"`
	TrackInfo  []struct {
		PlanDate     string `json:"plan_date"`
		ActualDate   string `json:"actual_date"`
		Event        string `json:"event"`
		Station      string `json:"station"`
		FlightNumber string `json:"flight_number"`
		Status       string `json:"status"`
		Piece        string `json:"piece"`
		Weight       string `json:"weight"`
	} `json:"track_info"`
	AirlineInfo struct {
		Name          string `json:"name"`
		Url           string `json:"url"`
		TrackUrl      string `json:"track_url"`
		TrackpageUrl2 string `json:"trackpage_url2"`
	} `json:"airline_info"`
}

type FlightDetails struct {
	DepartStation   string `json:"depart_station"`
	ArrivalStation  string `json:"arrival_station"`
	PlanDepartTime  string `json:"plan_depart_time"`
	DepartTime      string `json:"depart_time"`
	PlanArrivalTime string `json:"plan_arrival_time"`
	ArrivalTime     string `json:"arrival_time"`
}

type AirWaybillParams struct {
	AwbNumber string `json:"awb_number"`
}

func (client *Client) CreateAnAirWayBill(ctx context.Context, params AirWaybillParams) (*Response, error) {
	if params.AwbNumber == "" {
		return nil, errors.New(ErrMissingAwbNumber)
	}

	regexPattern := `^\d{3}[ -]?(\d{8})$`
	matched, err := regexp.MatchString(regexPattern, params.AwbNumber)
	if err != nil {
		return nil, err
	}

	if !matched {
		return nil, errors.New(ErrInvalidAirWaybillFormat)
	}

	var airWaybillItem AirWaybillItem
	response, err := client.sendApiRequest(ctx, http.MethodPost, "/awb", nil, params, &airWaybillItem)
	return response, err
}
