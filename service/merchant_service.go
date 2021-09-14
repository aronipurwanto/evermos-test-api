package service

import (
	"ahmadroni/test-evermos-api/model/web"
	"context"
)

type MerchantService interface {
	Create(ctx context.Context, request web.MerchantCreateRequest) web.MerchantResponse
	Update(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantResponse
	Delete(ctx context.Context, MerchantId int)
	FindById(ctx context.Context, MerchantId int) web.MerchantResponse
	FindAll(ctx context.Context) []web.MerchantResponse
}
