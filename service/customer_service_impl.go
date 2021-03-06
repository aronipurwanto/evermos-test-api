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

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(CustomerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: CustomerRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := domain.Customer{
		Name: request.Name,
	}

	customer = service.CustomerRepository.Save(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customer.Name = request.Name

	customer = service.CustomerRepository.Update(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Delete(ctx context.Context, CustomerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, CustomerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.Delete(ctx, tx, customer)
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, CustomerId int) web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, CustomerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) []web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CustomerRepository.FindAll(ctx, tx)

	return helper.ToCustomerResponses(categories)
}
