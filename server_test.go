package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPairDeviceHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/pair-device", nil)
	rec := httptest.NewRecorder()
	PairDeviceHandler(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("Expect 200 OK but get", rec.Code)
	}

	expected := `{"status":"active"}`
	if expected != rec.Body.String() {
		t.Errorf("Expected %q but got %q\n", expected, rec.Body.String())
	}
}
