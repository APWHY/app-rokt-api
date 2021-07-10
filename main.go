package main

import (
	"fmt"
	"net/http"

	"github.com/APWHY/app-rokt-api/autocomplete"
	"github.com/APWHY/app-rokt-api/nearby"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/nearby", nearby.GetNearby)
	http.HandleFunc("/autocomplete/applicant", autocomplete.GetApplicantSuggestion)
	http.HandleFunc("/autocomplete/address", autocomplete.GetAddressSuggestion)

	fmt.Print("Serving...\n")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
