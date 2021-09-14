package domain

type Product struct {
	Id			int
	MerchantId	int
	CategoryId	int
	Name 		string
	ImagePath	string
	Rating		float64
	Price		int
	Stock		int
}
