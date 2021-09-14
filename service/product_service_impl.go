package service

import (
	"ahmadroni/test-evermos-api/exception"
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/domain"
	"ahmadroni/test-evermos-api/model/web"
	"ahmadroni/test-evermos-api/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}


func NewProductService(ProductRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product := domain.Product{
		Name: request.Name,
	}

	Product = service.ProductRepository.Save(ctx, tx, Product)

	return helper.ToProductResponse(Product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	Product.Name = request.Name

	Product = service.ProductRepository.Update(ctx, tx, Product)

	return helper.ToProductResponse(Product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, ProductId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product, err := service.ProductRepository.FindById(ctx, tx, ProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, Product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, ProductId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product, err := service.ProductRepository.FindById(ctx, tx, ProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(Product)
}

func (service *ProductServiceImpl) FindByName(ctx context.Context, name string) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Products := service.ProductRepository.FindByName(ctx, tx, name)

	return helper.ToProductResponses(Products)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context, merchantId int) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Products := service.ProductRepository.FindAll(ctx, tx, merchantId)

	return helper.ToProductResponses(Products)
}
