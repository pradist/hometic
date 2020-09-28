package main

import (
	"bytes"
	"encoding/json"
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


	//orgin := createPairDevice
	//defer func() {
	//	createPairDevice = orgin
	//}()
	//createPairDevice = func(p Pair) error {
	//	return nil
	//}

	handler := &PairDeviceHandler{func(p Pair) error {
		return nil
	}}
	handler.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("Expect 200 OK but get", rec.Code)
	}

	expected := `{"status":"active"}`
	if expected != rec.Body.String() {
		t.Errorf("Expected %q but got %q\n", expected, rec.Body.String())
	}
}
