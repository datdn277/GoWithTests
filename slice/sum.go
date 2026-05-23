package slice

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAllTails(arr ...[]int) []int {
	var res []int
	for _, v := range arr {
		size := len(v)
		if size > 0 {
			res = append(res, Sum(v[1:]))
		} else {
			res = append(res, 0)
		}
	}
	return res
}
