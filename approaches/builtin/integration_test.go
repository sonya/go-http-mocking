// +build integration

package builtin_test

import (
	"testing"

	"github.com/sonya/go-http-mocking/builtin"
)

func TestGetFixedValue(t *testing.T) {
	value, err := builtin.GetFixedValue("http://localhost:8080")
	if err != nil {
		t.Error(err)
	}

	if value != "actualvalue" {
		t.Errorf("Expected 'actualvalue', got %s", value)
	}
}
