package autocomplete

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/APWHY/app-rokt-api/util"
)

func GetApplicantSuggestion(w http.ResponseWriter, r *http.Request) {
	term, err := util.PullFromQuery("term", r.URL.Query())
	fmt.Println(r.URL.Query()["term"][0])
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	suggestedPermits := []*util.Permit{}
	// using a trie would be better but if we use a db with a text_pattern_ops index or some equivalent
	for _, p := range util.Permits {
		if IsInitialSubstring(term, p.Applicant) {
			suggestedPermits = append(suggestedPermits, p)
		}
	}

	jsonResult, err := json.Marshal(suggestedPermits)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	w.Write(jsonResult)
}

func GetAddressSuggestion(w http.ResponseWriter, r *http.Request) {
	term, err := util.PullFromQuery("term", r.URL.Query())
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	suggestedPermits := []*util.Permit{}
	for _, p := range util.Permits {
		if IsSubstringOf(term, p.Address) {
			suggestedPermits = append(suggestedPermits, p)
		}
	}

	jsonResult, err := json.Marshal(suggestedPermits)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	w.Write(jsonResult)
}

func IsInitialSubstring(sub string, main string) bool {
	return strings.ToLower(main[:len(sub)]) == strings.ToLower(sub)
}

func IsSubstringOf(sub string, main string) bool {
	return strings.Contains(strings.ToLower(main), strings.ToLower(sub))
}
