package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0

	for i := 1; i < limit; i++ {
		for _, v := range divisors {
			if v != 0 && i%v == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}

// func SumMultiples(limit int, divisors ...int) int {

// 	var divs []int

// 	for _, d := range divisors {
// 		if d == 0 {
// 			continue
// 		}
// 		for i := d; i < limit; i += d {
// 			divs = append(divs, i)
// 		}
// 	}

// 	divs = unique(divs)

// 	sum := 0
// 	for _, d := range divs {
// 		sum += d
// 	}

// 	return sum
// }

// func unique(intSlice []int) []int {
// 	keys := make(map[int]bool, len(intSlice))
// 	list := []int{}
// 	for _, entry := range intSlice {
// 		if _, value := keys[entry]; !value {
// 			keys[entry] = true
// 			list = append(list, entry)
// 		}
// 	}
// 	return list
// }
