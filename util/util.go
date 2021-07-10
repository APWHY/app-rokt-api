package util

import (
	"fmt"
	"net/http"
	"net/url"
)

func PullFromQuery(field string, vs url.Values) (string, error) {
	if valueArr, ok := vs[field]; ok && len(valueArr[0]) > 0 {
		return valueArr[0], nil
	}
	return "", fmt.Errorf("Invalid or missing %s.", field)
}

func ErrorHandler(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Error: %s", err.Error())
}
