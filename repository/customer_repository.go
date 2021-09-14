package repository

import (
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
)

type MerchantRepository interface {
	Save(ctx context.Context, tx *sql.Tx, data domain.Merchant) domain.Merchant
	Update(ctx context.Context, tx *sql.Tx, data domain.Merchant) domain.Merchant
	Delete(ctx context.Context, tx *sql.Tx, data domain.Merchant)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Merchant, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Merchant
}
