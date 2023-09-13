package trackingmore

import (
	"context"
	"errors"
	"net/http"
)

type CreateTrackingParams struct {
	TrackingNumber             string `json:"tracking_number"`
	CourierCode                string `json:"courier_code"`
	OrderNumber                string `json:"order_number"`
	OriginCountryIso2          string `json:"origin_country_iso2"`
	DestinationCountryIso2     string `json:"destination_country_iso2"`
	CustomerName               string `json:"customer_name"`
	CustomerEmail              string `json:"customer_email"`
	CustomerSms                string `json:"customer_sms"`
	Title                      string `json:"title"`
	LogisticsChannel           string `json:"logistics_channel"`
	OrderId                    string `json:"order_id"`
	OrderDate                  string `json:"order_date"`
	TrackingCourierAccount     string `json:"tracking_courier_account"`
	TrackingPostalCode         string `json:"tracking_postal_code"`
	TrackingOriginCountry      string `json:"tracking_origin_country"`
	TrackingDestinationCountry string `json:"tracking_destination_country"`
	TrackingShipDate           string `json:"tracking_ship_date"`
	TrackingKey                string `json:"tracking_key"`
	Language                   string `json:"language"`
	Note                       string `json:"note"`
	AutoCorrect                string `json:"auto_correct"`
}

type Tracking struct {
	Id                         string `json:"id"`
	TrackingNumber             string `json:"tracking_number"`
	CourierCode                string `json:"courier_code"`
	OrderNumber                string `json:"order_number"`
	OrderDate                  string `json:"order_date"`
	CreatedAt                  string `json:"created_at"`
	UpdateAt                   string `json:"update_at"`
	DeliveryStatus             string `json:"delivery_status"`
	Archived                   string `json:"archived"`
	Updating                   bool   `json:"updating"`
	DestinationCountry         string `json:"destination_country"`
	DestinationState           string `json:"destination_state"`
	DestinationCity            string `json:"destination_city"`
	OriginCountry              string `json:"origin_country"`
	OriginState                string `json:"origin_state"`
	OriginCity                 string `json:"origin_city"`
	TrackingPostalCode         string `json:"tracking_postal_code"`
	TrackingShipDate           string `json:"tracking_ship_date"`
	TrackingDestinationCountry string `json:"tracking_destination_country"`
	TrackingOriginCountry      string `json:"tracking_origin_country"`
	TrackingKey                string `json:"tracking_key"`
	TrackingCourierAccount     string `json:"tracking_courier_account"`
	CustomerName               string `json:"customer_name"`
	CustomerEmail              string `json:"customer_email"`
	CustomerSms                string `json:"customer_sms"`
	OrderId                    string `json:"order_id"`
	Title                      string `json:"title"`
	LogisticsChannel           string `json:"logistics_channel"`
	Note                       string `json:"note"`
	SignedBy                   string `json:"signed_by"`
	ServiceCode                string `json:"service_code"`
	Weight                     string `json:"weight"`
	WeightKg                   string `json:"weight_kg"`
	ProductType                string `json:"product_type"`
	Pieces                     string `json:"pieces"`
	Dimension                  string `json:"dimension"`
	Previously                 string `json:"previously"`
	DestinationTrackNumber     string `json:"destination_track_number"`
	ExchangeNumber             string `json:"exchange_number"`
	ScheduledDeliveryDate      string `json:"scheduled_delivery_date"`
	ScheduledAddress           string `json:"scheduled_address"`
	Substatus                  string `json:"substatus"`
	StatusInfo                 string `json:"status_info"`
	LatestEvent                string `json:"latest_event"`
	LatestCheckpointTime       string `json:"latest_checkpoint_time"`
	TransitTime                int    `json:"transit_time"`
	OriginInfo                 struct {
		CourierCode     string `json:"courier_code"`
		CourierPhone    string `json:"courier_phone"`
		Weblink         string `json:"weblink"`
		ReferenceNumber string `json:"reference_number"`
		MilestoneDate   struct {
			InforeceivedDate   string `json:"inforeceived_date"`
			PickupDate         string `json:"pickup_date"`
			OutfordeliveryDate string `json:"outfordelivery_date"`
			DeliveryDate       string `json:"delivery_date"`
			ReturningDate      string `json:"returning_date"`
			ReturnedDate       string `json:"returned_date"`
		} `json:"milestone_date"`
		PickupDateLagacy             string `json:"pickup_date (Lagacy)"`
		DepartedAirportDateLagacy    string `json:"departed_airport_date (Lagacy)"`
		ArrivedAbroadDateLagacy      string `json:"arrived_abroad_date (Lagacy)"`
		CustomsReceivedDateLagacy    string `json:"customs_received_date (Lagacy)"`
		ArrivedDestinationDateLagacy string `json:"arrived_destination_date (Lagacy)"`
		Trackinfo                    []struct {
			CheckpointDate              string `json:"checkpoint_date"`
			CheckpointDeliveryStatus    string `json:"checkpoint_delivery_status"`
			CheckpointDeliverySubstatus string `json:"checkpoint_delivery_substatus"`
			TrackingDetail              string `json:"tracking_detail"`
			Location                    string `json:"location"`
			CountryIso2                 string `json:"country_iso2"`
			State                       string `json:"state"`
			City                        string `json:"city"`
			Zip                         string `json:"zip"`
			RawStatus                   string `json:"raw_status"`
		} `json:"trackinfo"`
	} `json:"origin_info"`
	DestinationInfo struct {
		CourierCode     string `json:"courier_code"`
		CourierPhone    string `json:"courier_phone"`
		Weblink         string `json:"weblink"`
		ReferenceNumber string `json:"reference_number"`
		MilestoneDate   struct {
			InforeceivedDate   string `json:"inforeceived_date"`
			PickupDate         string `json:"pickup_date"`
			OutfordeliveryDate string `json:"outfordelivery_date"`
			DeliveryDate       string `json:"delivery_date"`
			ReturningDate      string `json:"returning_date"`
			ReturnedDate       string `json:"returned_date"`
		} `json:"milestone_date"`
		PickupDate             string `json:"pickup_date"`
		DepartedAirportDate    string `json:"departed_airport_date"`
		ArrivedAbroadDate      string `json:"arrived_abroad_date"`
		CustomsReceivedDate    string `json:"customs_received_date"`
		ArrivedDestinationDate string `json:"arrived_destination_date"`
		Trackinfo              []struct {
			CheckpointDate              string `json:"checkpoint_date"`
			CheckpointDeliveryStatus    string `json:"checkpoint_delivery_status"`
			CheckpointDeliverySubstatus string `json:"checkpoint_delivery_substatus"`
			TrackingDetail              string `json:"tracking_detail"`
			Location                    string `json:"location"`
			CountryIso2                 string `json:"country_iso2"`
			State                       string `json:"state"`
			City                        string `json:"city"`
			Zip                         string `json:"zip"`
			RawStatus                   string `json:"raw_status"`
		} `json:"trackinfo"`
	} `json:"destination_info"`
}

type GetTrackingResultsParams struct {
	ArchivedStatus  string `json:"archived_status" url:"archived_status"`
	CourierCode     string `json:"courier_code" url:"courier_code"`
	CreatedDateMax  string `json:"created_date_max" url:"created_date_max"`
	CreatedDateMin  string `json:"created_date_min" url:"created_date_min"`
	DeliveryStatus  string `json:"delivery_status" url:"delivery_status"`
	ItemsAmount     int    `json:"items_amount" url:"items_amount"`
	Lang            string `json:"lang" url:"lang"`
	PagesAmount     int    `json:"pages_amount" url:"pages_amount"`
	TrackingNumbers string `json:"tracking_numbers" url:"tracking_numbers"`
	UpdatedDateMax  string `json:"updated_date_max" url:"updated_date_max"`
	UpdatedDateMin  string `json:"updated_date_min" url:"updated_date_min"`
}

type BatchResults struct {
	Success []BatchItem `json:"success"`
	Error   []BatchItem `json:"error"`
}

type BatchItem struct {
	Id             string `json:"id"`
	TrackingNumber string `json:"tracking_number"`
	CourierCode    string `json:"courier_code"`
	ErrorCode      int    `json:"errorCode,omitempty"`
	ErrorMessage   string `json:"errorMessage,omitempty"`
}

type UpdateTrackingParams struct {
	CourierCodeNew             string `json:"courier_code_new"`
	OrderNumber                string `json:"order_number"`
	OriginCountryIso2          string `json:"origin_country_iso2"`
	DestinationCountryIso2     string `json:"destination_country_iso2"`
	CustomerName               string `json:"customer_name"`
	CustomerEmail              string `json:"customer_email"`
	CustomerSms                string `json:"customer_sms"`
	Title                      string `json:"title"`
	LogisticsChannel           string `json:"logistics_channel"`
	OrderId                    string `json:"order_id"`
	OrderDate                  string `json:"order_date"`
	TrackingCourierAccount     string `json:"tracking_courier_account"`
	TrackingPostalCode         string `json:"tracking_postal_code"`
	TrackingOriginCountry      string `json:"tracking_origin_country"`
	TrackingDestinationCountry string `json:"tracking_destination_country"`
	TrackingShipDate           string `json:"tracking_ship_date"`
	TrackingKey                string `json:"tracking_key"`
	Language                   string `json:"language"`
	Note                       string `json:"note"`
}
type UpdateAfterResult struct {
	Id                         string `json:"id"`
	TrackingNumber             string `json:"tracking_number"`
	CourierCode                string `json:"courier_code"`
	CourierCodeNew             string `json:"courier_code_new"`
	OrderNumber                string `json:"order_number"`
	OrderDate                  string `json:"order_date"`
	OrderId                    string `json:"order_id"`
	DestinationCountry         string `json:"destination_country"`
	OriginCountry              string `json:"origin_country"`
	TrackingPostalCode         string `json:"tracking_postal_code"`
	TrackingShipDate           string `json:"tracking_ship_date"`
	TrackingCourierAccount     string `json:"tracking_courier_account"`
	TrackingDestinationCountry string `json:"tracking_destination_country"`
	TrackingOriginCountry      string `json:"tracking_origin_country"`
	TrackingKey                string `json:"tracking_key"`
	CustomerName               string `json:"customer_name"`
	CustomerEmail              string `json:"customer_email"`
	CustomerSms                string `json:"customer_sms"`
	Title                      string `json:"title"`
	LogisticsChannel           string `json:"logistics_channel"`
	Note                       string `json:"note"`
}

func (client *Client) CreateTracking(ctx context.Context, params CreateTrackingParams) (*Response, error) {
	if params.TrackingNumber == "" {
		return nil, errors.New(ErrMissingTrackingNumber)
	}
	if params.CourierCode == "" {
		return nil, errors.New(ErrMissingCourierCode)
	}
	var tracking Tracking
	response, err := client.sendApiRequest(ctx, http.MethodPost, "/trackings/create", nil, params, &tracking)
	return response, err
}

func (client *Client) GetTrackingResults(ctx context.Context, params GetTrackingResultsParams) (*Response, error) {
	var trackings []Tracking
	response, err := client.sendApiRequest(ctx, http.MethodGet, "/trackings/get", params, nil, &trackings)
	return response, err
}

func (client *Client) BatchCreateTrackings(ctx context.Context, params []CreateTrackingParams) (*Response, error) {
	if len(params) > 40 {
		return nil, errors.New(ErrMaxTrackingNumbersExceeded)
	}
	for _, item := range params {
		if item.TrackingNumber == "" {
			return nil, errors.New(ErrMissingTrackingNumber)
		}
		if item.CourierCode == "" {
			return nil, errors.New(ErrMissingCourierCode)
		}
	}
	var batchResults BatchResults
	response, err := client.sendApiRequest(ctx, http.MethodPost, "/trackings/batch", nil, params, &batchResults)
	return response, err
}

func (client *Client) UpdateTrackingByID(ctx context.Context, idString string, params UpdateTrackingParams) (*Response, error) {
	if idString == "" {
		return nil, errors.New(ErrEmptyId)
	}
	var updateAfterItem UpdateAfterResult
	response, err := client.sendApiRequest(ctx, http.MethodPut, "/trackings/update/"+idString, nil, params, &updateAfterItem)
	return response, err
}

func (client *Client) DeleteTrackingByID(ctx context.Context, idString string) (*Response, error) {
	if idString == "" {
		return nil, errors.New(ErrEmptyId)
	}
	var tracking Tracking
	response, err := client.sendApiRequest(ctx, http.MethodDelete, "/trackings/delete/"+idString, nil, nil, &tracking)
	return response, err
}

func (client *Client) RetrackTrackingByID(ctx context.Context, idString string) (*Response, error) {
	if idString == "" {
		return nil, errors.New(ErrEmptyId)
	}
	var tracking Tracking
	response, err := client.sendApiRequest(ctx, http.MethodPost, "/trackings/retrack/"+idString, nil, nil, &tracking)
	return response, err
}
