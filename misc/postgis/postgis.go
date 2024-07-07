package postgis

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Lat float64
	Lon float64
}
type PostGisGeo struct {
	Srid  string
	Point Point
}

func NewFromPoint(point Point) *PostGisGeo {
	return &PostGisGeo{
		Srid:  "4326",
		Point: point,
	}
}

func NewFromEwkt(raw string) (*PostGisGeo, error) {
	// sample: SRID=4326;POINT(12 121)
	split := strings.Split(raw, ";")
	if len(split) != 2 {
		return nil, errors.New("ewkt is not valid")
	}
	srid := strings.Split(split[0], "=")[1]
	point, _ := strings.CutPrefix(split[1], "POINT")
	lonlat := strings.TrimPrefix(point, "(")
	lonlat = strings.TrimSuffix(lonlat, ")")
	lonLatSplit := strings.Split(lonlat, " ")
	lon, err := strconv.ParseFloat(lonLatSplit[0], 64)
	if err != nil {
		return nil, err
	}
	lat, err := strconv.ParseFloat(lonLatSplit[1], 64)
	if err != nil {
		return nil, err
	}

	return &PostGisGeo{
		Srid: srid,
		Point: Point{
			Lon: lon,
			Lat: lat,
		},
	}, nil
}

func (p PostGisGeo) Ewkt() string {
	return fmt.Sprintf("SRID=%s;POINT(%f %f)", p.Srid, p.Point.Lon, p.Point.Lat)
}
