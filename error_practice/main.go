package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(f float64) (float64, error) {
	if f < 0.0 {
		return 0, ErrNegativeSqrt(f)
	}
	return 0, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	//val, err := Sqrt(-2)
	fmt.Println(Sqrt(-2))
}
