package ex06

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountEx03(x uint64) int {
	var acc int
	var i uint
	for i = 0; i < 8; i++ {
		acc += int(pc[byte(x>>(i*8))])
	}
	return acc
}

func PopCountEx04(x uint64) int {
	var acc int
	var current uint64
	var i uint
	for i = 0; i < 64; i++ {
		current = x >> i
		acc += int(current & 1)
	}
	return acc
}

func PopCountEx05(x uint64) (cnt int) {
	for x != 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}
