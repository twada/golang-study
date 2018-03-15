package lconv

import "fmt"

type Foot float64
type Meter float64
 
func (f Foot) String() string { return fmt.Sprintf("%g'", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func MToF(m Meter) Foot { return Foot(m * 3.2808) }
func FToM(f Foot) Meter { return Meter(f * 0.3048) }
