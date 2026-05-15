package darts

func Score(x, y float64) int {

	score := 0

	sqDistance := x*x + y*y

	switch {
	case sqDistance <= 1:
		score = 10
	case sqDistance <= 25:
		score = 5
	case sqDistance <= 100:
		score = 1
	}

	// if sqDistance <= 1 {
	// 	score = 10
	// } else if sqDistance <= 25 {
	// 	score = 5
	// } else if sqDistance <= 100 {
	// 	score = 1
	// }

	return score
}
