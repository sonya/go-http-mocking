// +build integration

package usinggock_test

import (
	"testing"

	"github.com/sonya/go-http-mocking/usinggock"
)

func TestGetFixedValue(t *testing.T) {
	value, err := usinggock.GetFixedValue("http://localhost:8080")
	if err != nil {
		t.Error(err)
	}

	if value != "actualvalue" {
		t.Errorf("Expected 'actualvalue', got %s", value)
	}
}
