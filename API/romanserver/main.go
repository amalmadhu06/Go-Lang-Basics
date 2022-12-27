package main

import (
	"api/romanNumerals"
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	http.HandleFunc("/", Sample)

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func Sample(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")

	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - not found"))
		} else {
			fmt.Fprintf(w, "%q",
				html.EscapeString(romanNumerals.Numerals[number]))

		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
	}

}
