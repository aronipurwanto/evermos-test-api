package service

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/domain"
	"ahmadroni/test-evermos-api/model/web"
	"ahmadroni/test-evermos-api/repository"
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type OrderServiceImpl struct {
	OrderRepository 		repository.OrderRepository
	OrderDetailRespository	repository.OrderDetailRepository
	ProductRepository 		repository.ProductRepository
	DB                 		*sql.DB
	Validate           		*validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, orderDetailRespository repository.OrderDetailRepository, productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		OrderDetailRespository: orderDetailRespository,
		ProductRepository: productRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	if len(request.OrderDetail) == 0 {
		helper.PanicIfError(errors.New("Product Order is empty"))
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	total := 0
	// check item
	for _, item := range request.OrderDetail{
		product, err := service.ProductRepository.FindById(ctx, tx ,item.ProductId)
		helper.PanicIfError(err)
		if product.Stock < item.Quantity {
			helper.PanicIfError(errors.New("Product Stock not enough"))
			total = 0
			break
		}

		total += item.Price * item.Quantity
	}

	order := domain.Order{
		CustomerId: request.CustomerId,
		Total: total,
		PaymentMethod: request.PaymentMethod,
		PaymentStatus: request.PaymentStatus,
		ShippingName: request.ShippingName,
	}

	order = service.OrderRepository.Save(ctx, tx, order)

	var orderDetailResponses []web.OrderDetailResponse
	for _, item := range request.OrderDetail {
		orderDetail := domain.OrderDetail{
			OrderId:    order.Id,
			ProductId:  item.ProductId,
			MerchantId: item.MerchantId,
			Price:      item.Price,
			Quantity:   item.Quantity,
			Amount:     item.Amount,
		}
		orderDetail = service.OrderDetailRespository.Save(ctx,tx,orderDetail)

		orderDetailResponse := web.OrderDetailResponse{
			Id:         orderDetail.Id,
			OrderId:    orderDetail.OrderId,
			ProductId:  orderDetail.ProductId,
			MerchantId: orderDetail.MerchantId,
			Price:      orderDetail.Price,
			Quantity:   orderDetail.Quantity,
			Amount:     orderDetail.Amount,
		}
		orderDetailResponses = append(orderDetailResponses, orderDetailResponse)
	}
	response := web.OrderResponse{
		Id:             order.Id,
		CustomerId:     order.CustomerId,
		Total:          order.Total,
		PaymentMethod:  order.PaymentMethod,
		PaymentStatus:  order.PaymentStatus,
		CreatedAt:      time.Time{},
		ConfirmAt:      time.Time{},
		ShippingName:   order.ShippingName,
		ShippingAt:     order.ShippingAt,
		ShippingStatus: order.ShippingStatus,
		OrderDetail:    orderDetailResponses,
	}
	return response
}
