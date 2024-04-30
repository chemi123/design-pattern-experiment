package logic

import (
	"design-pattern-experiment/pkg/utils"
)

var (
	max = utils.Max[int]
)

func AddBigger(num1, num2 int) int {
	biggerNum := max(num1, num2)
	return biggerNum + biggerNum
}
