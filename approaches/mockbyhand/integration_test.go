// +build integration

package mockbyhand_test

import (
	"testing"

	"github.com/sonya/go-http-mocking/mockbyhand"
)

// This test will never actually hit localhost:8080
// because the client has been overwritten by a
// unit test in the "global" package
func TestGetFixedValue(t *testing.T) {
	value, err := mockbyhand.GetFixedValue("http://localhost:8080")
	if err != nil {
		t.Error(err)
	}

	if value != "actualvalue" {
		t.Errorf("Expected 'actualvalue', got %s", value)
	}
}
