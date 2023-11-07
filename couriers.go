package trackingmore

import (
	"context"
	"errors"
	"net/http"
)

type Courier struct {
	CourierName            string      `json:"courier_name"`
	CourierCode            string      `json:"courier_code"`
	CourierCountryIso2     interface{} `json:"courier_country_iso2"`
	CourierUrl             string      `json:"courier_url"`
	CourierPhone           string      `json:"courier_phone"`
	CourierType            string      `json:"courier_type"`
	TrackingRequiredFields interface{} `json:"tracking_required_fields"`
	OptionalFields         interface{} `json:"optional_fields"`
	DefaultLanguage        string      `json:"default_language"`
	SupportLanguage        []string    `json:"support_language"`
	CourierLogo            string      `json:"courier_logo"`
}

type DetectParams struct {
	TrackingNumber string `json:"tracking_number"`
}

func (client *Client) GetAllCouriers(ctx context.Context) (*Response, error) {
	var couriers []Courier
	response, err := client.sendApiRequest(ctx, http.MethodGet, "/couriers/all", nil, nil, &couriers)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *Client) Detect(ctx context.Context, params DetectParams) (*Response, error) {
	if params.TrackingNumber == "" {
		return nil, errors.New(ErrMissingTrackingNumber)
	}
	var couriers []Courier
	response, err := client.sendApiRequest(ctx, http.MethodPost, "/couriers/detect", nil, params, &couriers)
	if err != nil {
		return nil, err
	}

	return response, nil
}
