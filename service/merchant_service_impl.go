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

	merchant := domain.Merchant{
		Name: request.Name,
	}

	merchant = service.MerchantRepository.Save(ctx, tx, merchant)

	return helper.ToMerchantResponse(merchant)
}

func (service *MerchantServiceImpl) Update(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	merchant, err := service.MerchantRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	merchant.Name = request.Name

	merchant = service.MerchantRepository.Update(ctx, tx, merchant)

	return helper.ToMerchantResponse(merchant)
}

func (service *MerchantServiceImpl) Delete(ctx context.Context, MerchantId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	merchant, err := service.MerchantRepository.FindById(ctx, tx, MerchantId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MerchantRepository.Delete(ctx, tx, merchant)
}

func (service *MerchantServiceImpl) FindById(ctx context.Context, MerchantId int) web.MerchantResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	merchant, err := service.MerchantRepository.FindById(ctx, tx, MerchantId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMerchantResponse(merchant)
}

func (service *MerchantServiceImpl) FindAll(ctx context.Context) []web.MerchantResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	merchants := service.MerchantRepository.FindAll(ctx, tx)

	return helper.ToMerchantResponses(merchants)
}
