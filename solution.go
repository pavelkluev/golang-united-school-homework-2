package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  
// Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"


type SidesCount int

const (
	SidesCircle SidesCount = iota
	_
	_
	SidesTriangle SidesCount = iota
	SidesSquare SidesCount = iota
)


func CalcSquare(sideLen float64, sidesNum SidesCount) float64 {
	switch sidesNum{
	case SidesCircle: 
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return math.Pow(sideLen, 2)* (math.Sqrt(float64(SidesTriangle)) / 4)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	}
	return 0
}
