package distance

import (
	"math"
)

type Point struct {
	latitude  float64
	longitude float64
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
