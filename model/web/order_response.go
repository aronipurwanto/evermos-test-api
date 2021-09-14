package web

type ProductResponse struct {
	Id 				int `json:"id"`
	MerchantId 		int `json:"merchant_id"`
	CategoryId 		int `json:"category_id"`
	Name 			string `json:"name"`
	ImagePath 		string `json:"image_path"`
	Price 			int `json:"price"`
	Stock 			int `json:"stock"`
	Rating	 		int `json:"rating"`
}
