package main

import (
	"fmt"
	"math"
)

type Point struct {
	latitude  float64
	longitude float64
}

func setPoint(latitude float64, longitude float64) *Point {
	return &Point{
		latitude:  latitude,
		longitude: longitude,
	}
}

const earthRadius = 6371

func (pointA *Point) Distance(pointB *Point) float64 {
	degreesLat := (pointB.latitude - pointA.latitude) * math.Pi / 180
	degreesLong := (pointB.longitude - pointA.longitude) * math.Pi / 180
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(pointA.latitude*math.Pi/180)*
			math.Cos(pointB.latitude*math.Pi/180)*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := earthRadius * c

	return d
}

func main() {

	Athens := setPoint(37.983972, 23.727806)
	Amsterdam := setPoint(52.366667, 4.9)
	Berlin := setPoint(52.516667, 13.388889)

	res := Amsterdam.Distance(Berlin)
	fmt.Printf("The distance from point A to point B is %.2f kilometers.\n", res)

	fmt.Println(Athens)
	fmt.Println(Amsterdam)
	fmt.Println(Berlin)

}
