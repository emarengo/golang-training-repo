package main

import (
	"testing"
)

func TestDistanceAthensToAmsterdam(t *testing.T) {
	Athens := setPoint(37.983972, 23.727806)
	Amsterdam := setPoint(52.366667, 4.9)

	result := Athens.Distance(Amsterdam)

	if result != 2163.2310285824487 {
		t.Error("Athens.Distance(Amsterdam) FAILED.")
	} else {
		t.Log("Athens.Distance(Amsterdam) SUCCEDED.")
	}
}

func TestDistanceAmsterdamToBerlin(t *testing.T) {
	Amsterdam := setPoint(52.366667, 4.9)
	Berlin := setPoint(52.516667, 13.388889)

	result := Amsterdam.Distance(Berlin)

	if result != 575.2949643958796 {
		t.Error("Amsterdam.Distance(Berlin) FAILED.")
	} else {
		t.Log("Amsterdam.Distance(Berlin) SUCCEDED.")
	}
}

func TestDistanceBerlinToAthens(t *testing.T) {
	Berlin := setPoint(52.516667, 13.388889)
	Athens := setPoint(37.983972, 23.727806)

	result := Berlin.Distance(Athens)

	if result != 1803.1087879059255 {
		t.Error("Berlin.Distance(Athens) FAILED.")
	} else {
		t.Log("Berlin.Distance(Athens) SUCCEDED.")
	}
}

var res float64

func BenchmarkDivision(b *testing.B) {

	// Any initialization code comes here
	var res1 float64
	Athens := setPoint(37.983972, 23.727806)
	Amsterdam := setPoint(52.366667, 4.9)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res1 = Athens.Distance(Amsterdam)
	}
	res = res1
}
