package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRomain(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/roman?n=10", nil)
	w := httptest.NewRecorder()
	toRomanHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Invalid response")
	}

	if w.Body.String() != "Roman Glyph is X!" {
		t.Fatalf("Unexpected response %s", w.Body.String())
	}
}

func TestRomainJSON(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/roman?n=10", nil)
	w := httptest.NewRecorder()
	toRomanHandlerJSON(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Invalid response")
	}

	resp := response{}
	err := json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Roman != "X" {
		t.Fatalf("Unexpected roman %s", resp.Roman)
	}
}
