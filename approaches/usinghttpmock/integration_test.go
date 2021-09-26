// +build integration

package usinghttpmock_test

import (
	"testing"

	"github.com/sonya/go-http-mocking/usinghttpmock"
)

func TestGetFixedValue(t *testing.T) {
	value, err := usinghttpmock.GetFixedValue("http://localhost:8080")
	if err != nil {
		t.Error(err)
	}

	if value != "actualvalue" {
		t.Errorf("Expected 'actualvalue', got %s", value)
	}
}
