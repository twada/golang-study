package popcount

func PopCount(x uint64) int {
	var acc int
	var current uint64
	var i uint
	for i = 0; i < 64; i++ {
		current = x >> i
		acc += int(current & 1)
	}
	return acc
}
