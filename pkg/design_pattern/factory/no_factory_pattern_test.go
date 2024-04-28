package factory

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_no_factory_pattern(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIProduct := NewMockIProduct(ctrl)
	mockIProduct.EXPECT().String().Return("MockProduct").AnyTimes()

	logicGeneratingProductDirectly := NewLogicGeneratingProductDirectly(mockIProduct)

	expected := "LogicGeneratingProductDirectly uses MockProduct"
	actual := logicGeneratingProductDirectly.Logic()

	if expected != actual {
		t.Errorf("Not match.\n expected=%s\n actual=%s\n", expected, actual)
	}
}
