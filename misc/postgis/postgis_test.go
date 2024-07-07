package postgis_test

import (
	"testing"

	"github.com/albugowy15/nearby-shops/misc/postgis"
)


func TestNewPostgisFromEwkt(t *testing.T) {
  testEwkt := "SRID=4326;POINT(123 45)"
  result, err := postgis.NewFromEwkt(testEwkt)
  if err != nil {
    t.Errorf("expected no error, got: %v", err)
  }

  if result.Srid != "4326" {
    t.Errorf("expected SRID to be %s, got %s", "4326", result.Srid)
  }
  if result.Point.Lon != 123.0 {
    t.Errorf("expected Point.Lon to be %f, got %f", 123.0, result.Point.Lon)
  }
  if result.Point.Lat != 45.0 {
    t.Errorf("expected Point.Lat to be %f, got %f", 45.0, result.Point.Lat)
  }
}

