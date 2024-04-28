package factory

type ICreator interface {
	NewProductWithFactory() IProductWithFactory
}

type Creator struct{}

func (creator *Creator) NewProductWithFactory() IProductWithFactory {
	return &ProductWithFactory{}
}

type IProductWithFactory interface {
	String() string
}

type ProductWithFactory struct{}

func (productWithFactory *ProductWithFactory) String() string {
	return "ProductWithFactory"
}

type LogicUsingProductWithFactory struct {
	ProductWithFactory IProductWithFactory
}

func NewLogicUsingProductWithFactory(creator ICreator) *LogicUsingProductWithFactory {
	return &LogicUsingProductWithFactory{
		ProductWithFactory: creator.NewProductWithFactory(),
	}
}

func (logicUsingProductWithFactor *LogicUsingProductWithFactory) Logic() string {
	return "LogicUsingProductWithFactory uses " + logicUsingProductWithFactor.ProductWithFactory.String()
}
