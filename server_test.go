package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPairDeviceHandler(t *testing.T) {

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(Pair{
		DeviceID: 1234,
		UserID:   4433,
	})

	req := httptest.NewRequest(http.MethodPost, "/pair-device", payload)
	rec := httptest.NewRecorder()

	create := func(p Pair) error {
		return nil
	}

	handler := CustomHandleFunc(PairDeviceHandler(CreatePairDeviceFunc(create)))
	handler.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("Expect 200 OK but get", rec.Code)
	}

	//expected := `{"status":"active"}`
	expected := fmt.Sprintf("%v\n", `{"status":"active"}`)
	if expected != rec.Body.String() {
		t.Errorf("Expected %q but got %q\n", expected, rec.Body.String())
	}
}