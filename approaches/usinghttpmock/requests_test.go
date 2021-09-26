package usinghttpmock

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetFixedValue(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://example.com/fixedvalue",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Accept") != "application/json" {
				t.Errorf("Expected Accept: application/json header, got: %s", req.Header.Get("Accept"))
			}
			resp, err := httpmock.NewJsonResponse(200, map[string]interface{}{
				"value": "fixed",
			})
			return resp, err
		},
	)

	value, _ := GetFixedValue("https://example.com")
	if value != "fixed" {
		t.Errorf("Expected 'fixed', got %s", value)
	}
}
