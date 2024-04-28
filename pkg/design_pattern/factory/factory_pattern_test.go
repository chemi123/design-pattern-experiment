package factory

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_factory_pattern(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockICreator := NewMockICreator(ctrl)
	mockIProductWithFactory := NewMockIProductWithFactory(ctrl)

	mockICreator.EXPECT().NewProductWithFactory().Return(mockIProductWithFactory).AnyTimes()
	mockIProductWithFactory.EXPECT().String().Return("MockProductWithFactory").AnyTimes()

	logicUsingProductWithFactory := NewLogicUsingProductWithFactory(mockICreator)

	expected := "LogicUsingProductWithFactory uses MockProductWithFactory"
	actual := logicUsingProductWithFactory.Logic()

	if expected != actual {
		t.Errorf("Not match.\n expected=%s\n actual=%s\n", expected, actual)
	}
}
