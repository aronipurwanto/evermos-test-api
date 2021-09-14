package service

import (
	"ahmadroni/test-evermos-api/model/web"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, ProductId int)
	FindById(ctx context.Context, ProductId int) web.ProductResponse
	FindByName(ctx context.Context, name string) []web.ProductResponse
	FindAll(ctx context.Context, merchantId int) []web.ProductResponse
}
