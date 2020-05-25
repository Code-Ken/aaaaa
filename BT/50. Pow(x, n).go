package BT


/*
快速幂 + 递归
*/

func myPow(x float64, n int) float64 {

	if n < 0 {
		x = 1 / x
		n = -n
	}

	return myFastPow(x, n)
}

func myFastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	half := myFastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

