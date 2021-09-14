package web

type MerchantUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
	Email 		string `validate:"required,min=1,max=100" json:"email"`
	Address 	string `validate:"required,min=1,max=100" json:"address"`
	Rating	 	float64 `validate:"required,min=1,max=100" json:"rating"`
}
