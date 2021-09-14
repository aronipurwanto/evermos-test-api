package service

import (
	"ahmadroni/test-evermos-api/model/web"
	"context"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
}
