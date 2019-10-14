package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSpecialCountRoute(t *testing.T) {
	router := setupRouter()

	text := "aaabaaa"

	var buf bytes.Buffer
	body := map[string]interface{}{
		"data": map[string]interface{}{
			"attributes": map[string]interface{}{
				"text": text,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		t.Error("Unable to encode test request body.")
	}

	reader := bytes.NewReader(buf.Bytes())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/special/count", reader)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Error("Expected status code 200 got", w.Code)
	}

	expected := "{\"data\":{\"attributes\":{\"ct\":16}}}"
	res := w.Body.String()
	if res != expected {
		t.Error("Expected response", expected, "got", res)
	}

}
