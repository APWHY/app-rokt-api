package util

// this is the definition of the permit object. It is odd to have it in a folder called util but this would normally be in the db schema

import (
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/umahmood/haversine"
)

// Not including all fields as they don't really seem that relevant
type Permit struct {
	LocationId string  `csv:"location_id"`
	Applicant  string  `csv:"Applicant"`
	Address    string  `csv:"Address"`
	Long       float64 `csv:"Longitude"`
	Lat        float64 `csv:"Latitude"`
}

var Permits []*Permit

func init() {
	// the relative address of data.csv is different during go.test, so we don't load it then
	if !strings.HasSuffix(os.Args[0], ".test") {
		dataFile, err := os.OpenFile("data.csv", os.O_RDONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer dataFile.Close()
		if err := gocsv.UnmarshalFile(dataFile, &Permits); err != nil { // Load permits from file
			panic(err)
		}
	}

}

// DistanceFrom returns the distwance between a coordinate and the Permit location in metres
func (p *Permit) DistanceFrom(long float64, lat float64) int {
	_, km := haversine.Distance(haversine.Coord{Lon: p.Long, Lat: p.Lat}, haversine.Coord{Lon: long, Lat: lat})
	return int(km * 1000) // we want to truncate instead of round
}
