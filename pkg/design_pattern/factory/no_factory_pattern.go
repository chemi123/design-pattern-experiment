package factory

type IProductWithoutFactory interface {
	String() string
}

type ProductWithoutFactory struct{}

func NewProductWithoutFactory() IProductWithoutFactory {
	return &ProductWithoutFactory{}
}

func (productWithoutFactory *ProductWithoutFactory) String() string {
	return "ProductWithoutFactory"
}

type LogicUsingProductWithoutFactory struct {
	ProductWithoutFactory IProductWithoutFactory
}

func NewLogicUsingProductWithoutFactory(iProductWithoutFactory IProductWithoutFactory) *LogicUsingProductWithoutFactory {
	return &LogicUsingProductWithoutFactory{
		ProductWithoutFactory: iProductWithoutFactory,
	}
}

func (logicUsingProductWithoutFactory *LogicUsingProductWithoutFactory) Logic() string {
	return "LogicUsingProductWithoutFactory uses " + logicUsingProductWithoutFactory.ProductWithoutFactory.String()
}
