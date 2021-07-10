package nearby

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/APWHY/app-rokt-api/util"
)

func GetNearby(w http.ResponseWriter, r *http.Request) {

	long, lat, radius, err := parseArgs(r)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	fmt.Fprintf(w, "Hello nearby, %s!\n", r.URL.Path[1:])
	fmt.Fprintf(w, "long:%f, lat:%f, radius:%d!", long, lat, radius) // assuming the default 6 decimal places is enough (which it should be)
}

func parseArgs(r *http.Request) (float64, float64, int, error) {
	q := r.URL.Query()
	longStr, err := util.PullFromQuery("long", q)
	if err != nil {
		return 0, 0, 0, err
	}
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {

		return 0, 0, 0, fmt.Errorf("Error: %s is not a valid value for longitude", longStr)
	}
	latStr, err := util.PullFromQuery("lat", q)
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {

		return 0, 0, 0, fmt.Errorf("Error: %s is not a valid value for latitude", latStr)
	}
	if err != nil {
		return 0, 0, 0, err
	}
	radius := 0
	if _, ok := q["radius"]; ok {
		radiusStr, err := util.PullFromQuery("radius", q)
		if err != nil {
			return 0, 0, 0, err
		}
		radius, err = strconv.Atoi(radiusStr)
		if err != nil {

			return 0, 0, 0, fmt.Errorf("Error: %s is not a valid value for radius", radiusStr)
		}
	}
	return long, lat, radius, nil
}
