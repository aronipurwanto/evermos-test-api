package repository

import (
	"ahmadroni/test-evermos-api/helper"
	"ahmadroni/test-evermos-api/model/domain"
	"context"
	"database/sql"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "insert into orders(customer_id, total,payment_method, payment_status, shipping_name, shipping_status) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, order.CustomerId, order.PaymentMethod, "NEW", order.ShippingName, "NEW")
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.Id = int(id)
	return order
}