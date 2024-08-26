package entity

type ProductID string

// domainの文脈において一意に定まるオブジェクト

type Product struct {
	ProductID   ProductID
	ProductName string
}
