package repository

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type MerchantRepositoryImpl struct {
}

func NewMerchantRepository() MerchantRepository {
	return &MerchantRepositoryImpl{}
}

func (repository *MerchantRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, merchant domain.Merchant) domain.Merchant {
	SQL := "insert into merchant(name, email, address, rating) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, merchant.Name, merchant.Email, merchant.Address, merchant.Rating)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	merchant.Id = int(id)
	return merchant
}

func (repository *MerchantRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, merchant domain.Merchant) domain.Merchant {
	SQL := "update merchant set name = ?, email = ?, address = ?, rating = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, merchant.Name, merchant.Email, merchant.Address, merchant.Rating, merchant.Id)
	helper.PanicIfError(err)

	return merchant
}

func (repository *MerchantRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, merchant domain.Merchant) {
	SQL := "delete from merchant where id = ?"
	_, err := tx.ExecContext(ctx, SQL, merchant.Id)
	helper.PanicIfError(err)
}

func (repository *MerchantRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, merchantId int) (domain.Merchant, error) {
	SQL := "select id, name, email, address, rating from merchant where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, merchantId)
	helper.PanicIfError(err)
	defer rows.Close()

	merchant := domain.Merchant{}
	if rows.Next() {
		err := rows.Scan(&merchant.Id, &merchant.Name, &merchant.Email, &merchant.Address, &merchant.Rating)
		helper.PanicIfError(err)
		return merchant, nil
	} else {
		return merchant, errors.New("merchant is not found")
	}
}

func (repository *MerchantRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Merchant {
	SQL := "select id, name, email, address, rating from merchant"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var merchants []domain.Merchant
	for rows.Next() {
		merchant := domain.Merchant{}
		err := rows.Scan(&merchant.Id, &merchant.Name, &merchant.Email, &merchant.Address, &merchant.Rating)
		helper.PanicIfError(err)
		merchants = append(merchants, merchant)
	}
	return merchants
}
