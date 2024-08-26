package query

import (
	"context"
	"design-pattern-experiment/pkg/design_pattern/aspect_pattern/domain/model/entity"
)

// IProduceQueryServiceはrefactorを割愛
type IProductQueryService interface {
	// Query Services(Read)
	GetFeaturedProducts(ctx context.Context) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, productID entity.ProductID) (*entity.Product, error)
}
