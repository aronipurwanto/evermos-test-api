package repository

import (
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
)

type OrderDetailRepository interface {
	Save(ctx context.Context, tx *sql.Tx, data domain.OrderDetail) domain.OrderDetail
}
