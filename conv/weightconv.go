package conv

import "fmt"

type Lb float64
type Kg float64

const KgInLb = 0.4535

func (l Lb) String() string { return fmt.Sprintf("%g Lb", l) }
func (k Kg) String() string { return fmt.Sprintf("%g Kg", k) }
