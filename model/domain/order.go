package domain

import "time"

type Order struct {
	Id				int
	CustomerId 		int
	Total	 		int
	PaymentMethod	string
	PaymentStatus	string
	CreatedAt		time.Time
	ConfirmAt		time.Time
	ShippingName	string
	ShippingAt		string
	ShippingStatus	string
}
