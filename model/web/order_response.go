package web

import "time"

type OrderResponse struct {
	Id				int `json:"id"`
	CustomerId 		int `json:"customer_id"`
	Total	 		int	`json:"total"`
	PaymentMethod	string `json:"payment_method"`
	PaymentStatus	string	`json:"payment_status"`
	CreatedAt		time.Time `json:"created_at"`
	ConfirmAt		time.Time `json:"confirm_at"`
	ShippingName	string `json:"shipping_name"`
	ShippingAt		string `json:"shipping_at"`
	ShippingStatus	string `json:"shipping_status"`
	OrderDetail 	[]OrderDetailResponse `json:"order_detail"`
}

type OrderDetailResponse struct {
	Id 			int `json:"id"`
	OrderId		int `json:"order_id"`
	ProductId	int `json:"product_id"`
	MerchantId	int `json:"merchant_id"`
	Price		int `json:"price"`
	Quantity	int `json:"quantity"`
	Amount 		int `json:"amount"`
}