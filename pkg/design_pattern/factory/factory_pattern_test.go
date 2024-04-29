package factory

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_factory_pattern(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockICreator := NewMockICreator(ctrl)
	mockIProduct := NewMockIProduct(ctrl)

	mockProductType := ProductType("MockProductType")

	mockICreator.EXPECT().NewProduct(mockProductType).Return(mockIProduct).AnyTimes()
	mockIProduct.EXPECT().String().Return("MockProduct").AnyTimes()

	logicGeneratingProductViaFactory := NewLogicGeneratingProductViaFactory(mockICreator, mockProductType)

	expected := "LogicGeneratingProductViaFactory uses MockProduct"
	actual := logicGeneratingProductViaFactory.Logic()

	if expected != actual {
		t.Errorf("Not match.\n expected=%s\n actual=%s\n", expected, actual)
	}
}
