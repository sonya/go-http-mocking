package builtin

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFixedValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/fixedvalue" {
			t.Errorf("Expected to request '/fixedvalue', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"value":"fixed"}`))
	}))
	defer server.Close()

	value, _ := GetFixedValue(server.URL)
	if value != "fixed" {
		t.Errorf("Expected 'fixed', got %s", value)
	}
}
