package controllers

import (
	"testing"
)

func TestAuthorize_HandleAuthorizeTestCard(t *testing.T) {
	controller := Controller{}

	err := controller.handleAuthorizeTestCard("4000000000000119")
	if err == nil {
		t.Errorf("Expected authorize failure")
	}
}
