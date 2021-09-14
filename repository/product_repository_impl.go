package repository

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product {
	SQL := "insert into product(merchant_id, category_id, name, image_path, rating, price, stock) values (?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, Product.MerchantId, Product.CategoryId, Product.Name, Product.ImagePath, Product.Rating, Product.Price, Product.Stock)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	Product.Id = int(id)
	return Product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product {
	SQL := "update product set merchant_id=?, category_id=?, name=?, image_path=?, rating=?, price=?, stock=? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, Product.MerchantId, Product.CategoryId, Product.Name, Product.ImagePath, Product.Rating, Product.Price, Product.Stock, Product.Id)
	helper.PanicIfError(err)

	return Product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, Product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, Product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Product, error) {
	SQL := "select id, merchant_id, category_id, name, image_path, rating, price, stock from product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	Product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&Product.Id, &Product.MerchantId, &Product.CategoryId, &Product.Name, &Product.ImagePath, &Product.Rating, &Product.Price, &Product.Stock)
		helper.PanicIfError(err)
		return Product, nil
	} else {
		return Product, errors.New("Product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) []domain.Product {
	SQL := "select id, merchant_id, category_id, name, image_path, rating, price, stock from product where name like '%'?'%'"
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)
	defer rows.Close()

	var Products []domain.Product
	for rows.Next() {
		Product := domain.Product{}
		err := rows.Scan(&Product.Id, &Product.MerchantId, &Product.CategoryId, &Product.Name, &Product.ImagePath, &Product.Rating, &Product.Price, &Product.Stock)
		helper.PanicIfError(err)
		Products = append(Products, Product)
	}
	return Products
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, merchantId int) []domain.Product {
	SQL := "select id, merchant_id, category_id, name, image_path, rating, price, stock from product where merchant_id=?"
	rows, err := tx.QueryContext(ctx, SQL, merchantId)
	helper.PanicIfError(err)
	defer rows.Close()

	var Products []domain.Product
	for rows.Next() {
		Product := domain.Product{}
		err := rows.Scan(&Product.Id, &Product.MerchantId, &Product.CategoryId, &Product.Name, &Product.ImagePath, &Product.Rating, &Product.Price, &Product.Stock)
		helper.PanicIfError(err)
		Products = append(Products, Product)
	}
	return Products
}