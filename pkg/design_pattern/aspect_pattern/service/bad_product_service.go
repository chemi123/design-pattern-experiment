package service

import (
	"context"
	"design-pattern-experiment/pkg/design_pattern/aspect_pattern/domain/model/entity"
)

// DI本で紹介されていたBadプラクティス。本当はもっとメソッドがあったがここではよりシンプルに少なくしている
// solid原則のsingle responsibility, interface separation, open closedに反している
type IBadProductService interface {
	// Query Services(Read)
	GetFeaturedProducts(ctx context.Context) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, productID entity.ProductID) (*entity.Product, error)

	// Command Services(Write)
	DeleteProduct(ctx context.Context, productID entity.ProductID) error
	InsertProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, product *entity.Product) error
}
