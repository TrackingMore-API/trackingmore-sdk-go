package trackingmore

import (
	"context"
	"testing"
)

func TestCreateAnAirWayBill(t *testing.T) {
	key := "you api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	validParams := AirWaybillParams{
		AwbNumber: "235-69030430",
	}

	response, err := client.CreateAnAirWayBill(context.Background(), validParams)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Error("Expected a response, got nil")
	}

	var _, ok = response.Data.(*AirWaybillItem)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

	invalidParams := AirWaybillParams{
		AwbNumber: "",
	}

	response, err = client.CreateAnAirWayBill(context.Background(), invalidParams)
	if err == nil || err.Error() != ErrMissingAwbNumber {
		t.Errorf("Expected ErrMissingAwbNumber error, got %v", err)
	}
	if response != nil {
		t.Error("Expected nil response, got a response")
	}

	invalidParams.AwbNumber = "12345"

	response, err = client.CreateAnAirWayBill(context.Background(), invalidParams)
	if err == nil || err.Error() != ErrInvalidAirWaybillFormat {
		t.Errorf("Expected ErrInvalidAirWaybillFormat error, got %v", err)
	}
	if response != nil {
		t.Error("Expected nil response, got a response")
	}

}
