package usinggock

import (
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestGetFixedValue(t *testing.T) {
	defer gock.Off()

	gock.New("https://example.com/fixedvalue").
		MatchHeader("Accept", "application/json").
		Get("/").
		Reply(200).
		JSON(map[string]string{"value": "fixed"})

	value, _ := GetFixedValue("https://example.com")
	if value != "fixed" {
		t.Errorf("Expected 'fixed', got %s", value)
	}
}
