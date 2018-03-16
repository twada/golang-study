package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/twada/golang-study/ch02/lconv"
	"github.com/twada/golang-study/ch02/tempconv"
)

func main() {
	values := os.Args[1:]
	if len(values) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			conv(input.Text())
		}
	} else {
		for _, arg := range values {
			conv(arg)
		}
	}
}

func conv(v string) {
	t, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	ft := lconv.Foot(t)
	m := lconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n",
		ft, lconv.FToM(ft), m, lconv.MToF(m))
}
