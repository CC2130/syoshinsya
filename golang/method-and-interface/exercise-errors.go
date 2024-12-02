package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can't Sqrt negative number: %v", float64(*e))
}

// if f is negative, retrun ErrNegativeSqrt
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, &ErrNegativeSqrt(f)
		err := ErrNegativeSqrt(f)
		return 0, &err
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - f) / (2 * z)
	}
	return z, nil
}

func main() {
	_, err := Sqrt(-4)
	if err != nil {
		fmt.Println(err)
	}
}
