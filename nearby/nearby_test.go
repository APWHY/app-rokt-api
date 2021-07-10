package nearby

import (
	"net/http"
	"net/url"
	"testing"
)

func TestParseArgsShouldPass(t *testing.T) {
	type tester struct {
		in       string
		exLong   float64
		exLat    float64
		exRadius int
	}
	ins := []*tester{
		{
			in:       "http://localhost:8080/nearby?long=-122.39006184632663&lat=37.722629217598346",
			exLong:   -122.39006184632663,
			exLat:    37.722629217598346,
			exRadius: 0,
		},
		{
			in:       "http://localhost:8080/nearby?long=-122.388369&lat=37.725498&radius=90",
			exLong:   -122.388369,
			exLat:    37.725498,
			exRadius: 90,
		},
	}

	for _, in := range ins {
		URL, _ := url.Parse(in.in)
		long, lat, radius, err := parseArgs(&http.Request{URL: URL})
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if long != in.exLong {
			t.Errorf("long is wrong (expect, got) (%f, %f)", in.exLong, long)
		}
		if lat != in.exLat {
			t.Errorf("lat is wrong (expect, got) (%f, %f)", in.exLat, lat)
		}
		if radius != in.exRadius {
			t.Errorf("radius is wrong (expect, got) (%d, %d)", in.exRadius, radius)
		}
	}

}
func TestParseArgsShouldFail(t *testing.T) {
	type tester struct {
		in string
	}
	ins := []*tester{
		{
			in: "http://localhost:8080/nearby?long=-122.39006184632663&lat=37.722629217598346&radius",
		},
		{
			in: "http://localhost:8080/nearby?long=-122.388369&lat=frog&radius=90",
		},
		{
			in: "http://localhost:8080/nearby?long=-222.388369&lat=0&radius=90",
		},
		{
			in: "http://localhost:8080/nearby?long=0&lat=100&radius=90",
		},
	}

	for _, in := range ins {
		URL, _ := url.Parse(in.in)
		_, _, _, err := parseArgs(&http.Request{URL: URL})
		if err == nil {
			t.Errorf("Expected an error to be returned with argument %s but there was none", in.in)
		}

	}

}
