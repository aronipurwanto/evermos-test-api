package repository

import (
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, data domain.Order) domain.Order
}
