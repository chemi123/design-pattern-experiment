package factory

func NewProduct() IProduct {
	return &StandardProduct{}
}

type LogicGeneratingProductDirectly struct {
	Product IProduct
}

func NewLogicGeneratingProductDirectly(product IProduct) *LogicGeneratingProductDirectly {
	return &LogicGeneratingProductDirectly{
		Product: product,
	}
}

func (logicGeneratingProductDirectly *LogicGeneratingProductDirectly) Logic() string {
	return "LogicGeneratingProductDirectly uses " + logicGeneratingProductDirectly.Product.String()
}
