package mockbyhand

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetFixedValue(t *testing.T) {
	Client = &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.Path != "/fixedvalue" {
				t.Errorf("Expected to request '/fixedvalue', got: %s", req.URL.Path)
			}
			if req.Header.Get("Accept") != "application/json" {
				t.Errorf("Expected Accept: application/json header, got: %s", req.Header.Get("Accept"))
			}

			responseBody := ioutil.NopCloser(bytes.NewReader([]byte(`{"value":"fixed"}`)))
			return &http.Response{
				StatusCode: 200,
				Body:       responseBody,
			}, nil
		},
	}

	value, _ := GetFixedValue("https://example.com")
	if value != "fixed" {
		t.Errorf("Expected 'fixed', got %s", value)
	}
}
