package factory

type IProduct interface {
	String() string
}

type StandardProduct struct{}

func (standardProduct *StandardProduct) String() string {
	return "StandardProduct"
}

type PremiumProduct struct{}

func (premiumProduct *PremiumProduct) String() string {
	return "PremiumProduct"
}
