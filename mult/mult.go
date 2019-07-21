package mult

import "../sum/sum"

func Mult(a int, b int) int {
	ans := a
	for i := 0; i < b-1; i++ {
		ans = sum.Sum(ans, a)
	}

	return ans;
}