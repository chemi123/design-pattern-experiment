package factory

import "fmt"

type ICreator interface {
	NewProduct(productType ProductType) IProduct
}

type Creator struct {
	NewProductFuncMap map[ProductType]NewProductFunc
}

func NewCreator() *Creator {
	return &Creator{
		map[ProductType]NewProductFunc{
			StandardProductKey: newStandardProduct,
			PremiumProductKey:  newPremiumProduct,
		},
	}
}

func (creator *Creator) NewProduct(productType ProductType) IProduct {
	newProductFunc, ok := creator.NewProductFuncMap[productType]
	if !ok {
		panic(fmt.Sprintf("%s is not set in ProductFuncMap", productType))
	}

	return newProductFunc()
}

type LogicGeneratingProductViaFactory struct {
	Product IProduct
}

func NewLogicGeneratingProductViaFactory(creator ICreator, productType ProductType) *LogicGeneratingProductViaFactory {
	return &LogicGeneratingProductViaFactory{
		Product: creator.NewProduct(productType),
	}
}

func (logicGeneratingProductViaFactory *LogicGeneratingProductViaFactory) Logic() string {
	return "LogicGeneratingProductViaFactory uses " + logicGeneratingProductViaFactory.Product.String()
}

func newStandardProduct() IProduct {
	return &StandardProduct{}
}

func newPremiumProduct() IProduct {
	return &PremiumProduct{}
}
