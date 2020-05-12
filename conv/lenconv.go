package conv

import "fmt"

type Ft float64
type Meter float64

const MetersInFt = 0.3048

func (f Ft) String() string    { return fmt.Sprintf("%g Ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }
