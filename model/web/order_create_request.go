package web

type OrderCreateRequest struct {
	CustomerId 		int `validate:"required" json:"customer_id"`
	Total	 		int `validate:"required" json:"total"`
	PaymentMethod	string `validate:"required" json:"payment_method"`
	PaymentStatus	string `validate:"required" json:"payment_status"`
	ShippingName	string `validate:"required" json:"shipping_name"`
	OrderDetail     []OrderDetailRequest `json:"order_detail"`
}

type OrderDetailRequest struct {
	ProductId	int `validate:"required" json:"product_id"`
	MerchantId	int `validate:"required" json:"merchant_id"`
	Price		int `validate:"required" json:"price"`
	Quantity	int `validate:"required" json:"quantity"`
	Amount 		int `validate:"required" json:"amount"`
}