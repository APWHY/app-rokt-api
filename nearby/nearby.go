package nearby

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/APWHY/app-rokt-api/util"
)

const MAX_LONG = 180
const MIN_LONG = -180
const MAX_LAT = 90
const MIN_LAT = -90

func GetNearby(w http.ResponseWriter, r *http.Request) {

	long, lat, radius, err := parseArgs(r)
	if err != nil {
		util.ErrorHandler(w, http.StatusBadRequest, err)
		return
	}
	nearbyPermits := []*util.Permit{}
	for _, p := range util.Permits {
		if p.DistanceFrom(long, lat) <= radius {
			nearbyPermits = append(nearbyPermits, p)
		}
	}

	jsonResult, err := json.Marshal(nearbyPermits)
	if err != nil {
		util.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(jsonResult)
}

func parseArgs(r *http.Request) (float64, float64, int, error) {
	q := r.URL.Query()
	longStr, err := util.PullFromQuery("long", q)
	if err != nil {
		return 0, 0, 0, err
	}
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil || long > MAX_LONG || long < MIN_LONG {
		return 0, 0, 0, fmt.Errorf("Error: %s is not a valid value for longitude", longStr)
	}
	latStr, err := util.PullFromQuery("lat", q)
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil || lat > MAX_LAT || lat < MIN_LAT {

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
