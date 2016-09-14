package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArabicHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/arabic?g=X", nil)
	w := httptest.NewRecorder()
	toArabicHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Invalid response")
	}

	resp := response{}
	err := json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Result != "10" {
		t.Fatalf("Unexpected arabic number %s", resp.Result)
	}
}

func TestRomainHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/roman?n=10", nil)
	w := httptest.NewRecorder()
	toRomanHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Invalid response")
	}

	resp := response{}
	err := json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Result != "X" {
		t.Fatalf("Unexpected roman %s", resp.Result)
	}
}
