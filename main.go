package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/nearby", GetNearby)

	fmt.Print("Serving...\n")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func GetNearby(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	longStr, err := pullFromQuery("long", q)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %s is not a valid value for longitude", longStr)
		return
	}
	latStr, err := pullFromQuery("lat", q)
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %s is not a valid value for latitude", latStr)
		return
	}
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	radius := 0
	if _, ok := q["radius"]; ok {
		radiusStr, err := pullFromQuery("radius", q)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		radius, err = strconv.Atoi(radiusStr)
		if err != nil {
			fmt.Fprintf(w, "Error: %s is not a valid value for radius", radiusStr)
			return
		}
	}
	fmt.Fprintf(w, "Hello nearby, %s!\n", r.URL.Path[1:])
	fmt.Fprintf(w, "long:%f, lat:%f, radius:%d!", long, lat, radius)
}

func pullFromQuery(field string, vs url.Values) (string, error) {
	if valueArr, ok := vs[field]; ok && len(valueArr[0]) > 1 {
		return valueArr[0], nil
	}
	return "", fmt.Errorf("Invalid or missing %s.", field)
}
