package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/derailed/imhotep/golabs/roman"
)

type response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
	URL    string `json:"url"`
}

func toRomanHandler(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	glyph := roman.ToRoman(number)
	resp := response{
		Status: http.StatusOK,
		Result: glyph,
		URL:    fmt.Sprintf("http://%s/arabic?g=%s", r.Host, glyph),
	}

	buff := new(bytes.Buffer)
	err = json.NewEncoder(buff).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, buff.String())
	log.Printf("[%d] %s", http.StatusOK, r.URL)
}

func toArabicHandler(w http.ResponseWriter, r *http.Request) {
	n := roman.ToArabic(r.URL.Query().Get("g"))
	resp := response{
		Status: http.StatusOK,
		Result: fmt.Sprintf("%d", n),
		URL:    fmt.Sprintf("http://%s/roman?n=%d", r.Host, n),
	}

	buff := new(bytes.Buffer)
	err := json.NewEncoder(buff).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, buff.String())
	log.Printf("[%d] %s", http.StatusOK, r.URL)
}

func noMatchHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("No matching routes! %s", r.URL), http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/roman", toRomanHandler)
	http.HandleFunc("/arabic", toArabicHandler)
	http.HandleFunc("/", noMatchHandler)
	http.ListenAndServe(":9000", nil)
}
