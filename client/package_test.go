package client

import (
	"exec/series"
	"testing"
	)


func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	//t.Log(series.square(2))
}