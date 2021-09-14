package service

import (
	"ahmadroni/test-evermos-api/model/web"
	"context"
)

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse
	Delete(ctx context.Context, CustomerId int)
	FindById(ctx context.Context, CustomerId int) web.CustomerResponse
	FindAll(ctx context.Context) []web.CustomerResponse
}
