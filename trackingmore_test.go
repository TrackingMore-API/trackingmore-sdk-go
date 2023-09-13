package trackingmore

import (
	"errors"
	"testing"
)

func TestNewClientWithValidAPIKey(t *testing.T) {
	apiKey := "you api key"
	client, err := NewClient(apiKey)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Error("Expected a non-nil client, got nil")
	}

	if client.apiKey != apiKey {
		t.Errorf("Expected API key %s, got %s", apiKey, client.apiKey)
	}
}

func TestNewClientWithEmptyAPIKey(t *testing.T) {
	apiKey := ""
	client, err := NewClient(apiKey)

	expectedError := errors.New(ErrEmptyAPIKey)
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if client != nil {
		t.Error("Expected a nil client, got non-nil client")
	}
}
