package factory

type ICreator interface {
	NewProduct() IProduct
}

type Creator struct{}

func (creator *Creator) NewProduct() IProduct {
	return &StandardProduct{}
}

type LogicGeneratingProductViaFactory struct {
	Product IProduct
}

func NewLogicGeneratingProductViaFactory(creator ICreator) *LogicGeneratingProductViaFactory {
	return &LogicGeneratingProductViaFactory{
		Product: creator.NewProduct(),
	}
}

func (logicGeneratingProductViaFactory *LogicGeneratingProductViaFactory) Logic() string {
	return "LogicGeneratingProductViaFactory uses " + logicGeneratingProductViaFactory.Product.String()
}
