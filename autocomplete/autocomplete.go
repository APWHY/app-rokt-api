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
		util.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	suggestedPermits := []*util.Permit{}
	// using a trie would be better but if we use a db with a text_pattern_ops index or some equivalent
	for _, p := range util.Permits {
		if isInitialSubstring(term, p.Applicant) {
			suggestedPermits = append(suggestedPermits, p)
		}
	}

	jsonResult, err := json.Marshal(suggestedPermits)
	if err != nil {
		util.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(jsonResult)
}

func GetAddressSuggestion(w http.ResponseWriter, r *http.Request) {
	term, err := util.PullFromQuery("term", r.URL.Query())
	if err != nil {
		util.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}

	suggestedPermits := []*util.Permit{}
	for _, p := range util.Permits {
		if isSubstringOf(term, p.Address) {
			suggestedPermits = append(suggestedPermits, p)
		}
	}

	jsonResult, err := json.Marshal(suggestedPermits)
	if err != nil {
		util.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(jsonResult)
}

func isInitialSubstring(sub string, main string) bool {
	return len(main) >= len(sub) && strings.ToLower(main[:len(sub)]) == strings.ToLower(sub)
}

func isSubstringOf(sub string, main string) bool {
	return strings.Contains(strings.ToLower(main), strings.ToLower(sub))
}
