package factory

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_no_factory_pattern(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIProductWithoutFactory := NewMockIProductWithoutFactory(ctrl)
	mockIProductWithoutFactory.EXPECT().String().Return("MockProductWithoutFactory").AnyTimes()

	logicUsingProductWithoutFactory := NewLogicUsingProductWithoutFactory(mockIProductWithoutFactory)

	expected := "LogicUsingProductWithoutFactory uses MockProductWithoutFactory"
	actual := logicUsingProductWithoutFactory.Logic()

	if expected != actual {
		t.Errorf("Not match.\n expected=%s\n actual=%s\n", expected, actual)
	}
}
