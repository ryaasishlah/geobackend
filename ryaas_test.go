package geobackend

import (
	"fmt"
	"testing"
)

var dbname = "gisyas"
var collname = "post"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{107.600136283516, -6.904398057194598},
				{107.60231617360034, -6.911348950466206},
				{107.60703926522939, -6.91519895280679},
				{107.61198593822624, -6.908005322704},
				{107.600136283516, -6.904398057194598},
			},
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{107.60443214379671, -6.905131269636627},
				{107.60444655183989, -6.905832179706749},
				{107.60496553082828, -6.905809511262689},
				{107.60499936141781, -6.905136688887282},
				{107.60443214379671, -6.905131269636627},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "gisyas", "post")
	coordinates := Point{
		Coordinates: []float64{
			107.60858561038731, -6.9135888696341965,
		},
	}
	datagedung := Near(mconn, "post", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gisyas")
	coordinates := Point{
		Coordinates: []float64{
			107.60473077068144, -6.905467750563176,
		},
	}
	datagedung := NearSphere(mconn, "post", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gisyas")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{107.6067794338868, -6.914148270960638},
			{107.60707235969664, -6.914294215762908},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
