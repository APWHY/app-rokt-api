package util

// this is the definition of the permit object. It is odd to have it in a folder called util but this would normally be in the db schema

import (
	"os"

	"github.com/gocarina/gocsv"
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
	dataFile, err := os.OpenFile("data.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()
	if err := gocsv.UnmarshalFile(dataFile, &Permits); err != nil { // Load permits from file
		panic(err)
	}
}
