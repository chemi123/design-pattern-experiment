package utils

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](num1, num2 T) T {
	if num1 > num2 {
		return num1
	}
	return num2
}
