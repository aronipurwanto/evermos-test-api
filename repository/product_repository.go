package repository

import (
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, data domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, data domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, data domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Product, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) []domain.Product
	FindAll(ctx context.Context, tx *sql.Tx, merchantId int) []domain.Product
}
