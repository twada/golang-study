package popcount

func PopCount(x uint64) (cnt int) {
	for (x != 0) {
		x = x & (x-1)
		cnt++
	}
	return cnt
}
