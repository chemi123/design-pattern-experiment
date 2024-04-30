package logic

import (
	"design-pattern-experiment/pkg/utils"
	"testing"
)

func Test_AddBigger(t *testing.T) {
	defer func() {
		max = utils.Max[int]
	}()
	max = func(num1, num2 int) int { return num1 }

	num1, num2 := 10, 5
	expected := num1 + num1
	actual := AddBigger(num1, num2)

	if expected != actual {
		t.Error("Not match")
	}
}
