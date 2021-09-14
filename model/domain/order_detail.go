package domain

type OrderDetail struct {
	Id 			int
	OrderId		int
	ProductId	int
	MerchantId	int
	Price		int
	Quantity	int
	Amount 		int
}
