package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:

//   CalcSquare(10.0, SidesTriangle)
//   CalcSquare(10.0, SidesSquare)
//   CalcSquare(10.0, SidesCircle)
type SidesNumber int

const (
	SidesCircle   SidesNumber = 0
	SidesTriangle SidesNumber = 3
	SidesSquare   SidesNumber = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {
	var s float64
	if sidesNum == 0 {
		s = math.Pi * math.Pow(sideLen, 2)
		return s
	} else if sidesNum == 3 {
		s = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
		return s
	} else if sidesNum == 4 {
		s = sideLen * sideLen
		return s
	} else {
		return 0
	}
}
