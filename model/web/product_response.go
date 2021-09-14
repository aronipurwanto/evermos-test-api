package web

type ProductUpdateRequest struct {
	Id 				int `validate:"required" json:"id"`
	MerchantId 		string `validate:"required,min=1,max=100" json:"merchant_id"`
	CategoryId 		string `validate:"required,min=1,max=100" json:"category_id"`
	Name 			string `validate:"required,min=1,max=100" json:"name"`
	ImagePath 		string `validate:"required,min=1,max=100" json:"image_path"`
	Price 			int `validate:"required,min=1,max=100" json:"price"`
	Stock 			int `validate:"required,min=1,max=100" json:"stock"`
	Rating	 		int `validate:"required,min=1,max=100" json:"rating"`
}
