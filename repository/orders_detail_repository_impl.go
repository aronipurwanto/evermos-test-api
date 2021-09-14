package repository

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
)

type OrderDetailRepositoryImpl struct {
}

func NewOrderDetailRepository() OrderDetailRepository {
	return &OrderDetailRepositoryImpl{}
}

func (repository *OrderDetailRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.OrderDetail) domain.OrderDetail {
	SQL := "insert into order_details(order_id, product_id,merchant_id, price, quantity, amount) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, order.OrderId, order.ProductId, order.MerchantId, order.Price, order.Quantity, order.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.Id = int(id)
	return order
}