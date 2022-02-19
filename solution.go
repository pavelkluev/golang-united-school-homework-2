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
	circle SidesCount = iota
	_
	_
	triangle SidesCount = iota
	square SidesCount = iota
)


func CalcSquare(sideLen float64, sidesNum SidesCount) float64 {
	switch sidesNum{
	case circle: 
		return math.Pi * math.Pow(sideLen, 2)
	case triangle:
		return math.Pow(sideLen, 2)* (math.Sqrt(float64(triangle)) / 4)
	case square:
		return math.Pow(sideLen, 2)
	}
	return 0
}
