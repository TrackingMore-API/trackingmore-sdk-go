package trackingmore

import (
	"context"
	"testing"
)

func TestGetAllCouriers(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	response, err := client.GetAllCouriers(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Error("Expected a response, got nil")
	}

	var _, ok = response.Data.(*[]Courier)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}
}

func TestDetect(t *testing.T) {
	key := "your api key"
	client, err := NewClient(key)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	validParams := DetectParams{
		TrackingNumber: "1234567890",
	}

	response, err := client.Detect(context.Background(), validParams)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Error("Expected a response, got nil")
	}

	var _, ok = response.Data.(*[]Courier)
	if response.Meta.Code == 200 && !ok {
		t.Error("Structure type conversion failed")
	}

	emptyParams := DetectParams{
		TrackingNumber: "",
	}

	response, err = client.Detect(context.Background(), emptyParams)
	if err == nil || err.Error() != ErrMissingTrackingNumber {
		t.Errorf("Expected ErrMissingTrackingNumber error, got %v", err)
	}
	if response != nil {
		t.Error("Expected nil response, got a response")
	}

}
