package G20200343030501

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return 0 == n & (n - 1)
}