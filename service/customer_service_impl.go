package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"ahmadroni/test-evermos-api/exception"
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/domain"
	"ahmadroni/test-evermos-api/model/web"
	"ahmadroni/test-evermos-api/repository"
)

type MerchantServiceImpl struct {
	MerchantRepository repository.MerchantRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewMerchantService(MerchantRepository repository.MerchantRepository, DB *sql.DB, validate *validator.Validate) MerchantService {
	return &MerchantServiceImpl{
		MerchantRepository: MerchantRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *MerchantServiceImpl) Create(ctx context.Context, request web.MerchantCreateRequest) web.MerchantResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Merchant := domain.Merchant{
		Name: request.Name,
	}

	Merchant = service.MerchantRepository.Save(ctx, tx, Merchant)

	return helper.ToMerchantResponse(Merchant)
}

func (service *MerchantServiceImpl) Update(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Merchant, err := service.MerchantRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	Merchant.Name = request.Name

	Merchant = service.MerchantRepository.Update(ctx, tx, Merchant)

	return helper.ToMerchantResponse(Merchant)
}

func (service *MerchantServiceImpl) Delete(ctx context.Context, MerchantId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Merchant, err := service.MerchantRepository.FindById(ctx, tx, MerchantId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MerchantRepository.Delete(ctx, tx, Merchant)
}

func (service *MerchantServiceImpl) FindById(ctx context.Context, MerchantId int) web.MerchantResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Merchant, err := service.MerchantRepository.FindById(ctx, tx, MerchantId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMerchantResponse(Merchant)
}

func (service *MerchantServiceImpl) FindAll(ctx context.Context) []web.MerchantResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.MerchantRepository.FindAll(ctx, tx)

	return helper.ToMerchantResponses(categories)
}
