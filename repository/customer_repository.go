package repository

import (
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, data domain.Customer) domain.Customer
	Update(ctx context.Context, tx *sql.Tx, data domain.Customer) domain.Customer
	Delete(ctx context.Context, tx *sql.Tx, data domain.Customer)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
}
