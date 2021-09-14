package web

type MerchantCreateRequest struct {
	Name 		string `validate:"required,min=1,max=100" json:"name"`
	Email 		string `validate:"required,min=1,max=100" json:"email"`
	Address 	string `validate:"required,min=1,max=100" json:"address"`
	Rating	 	float64 `validate:"required,min=1,max=100" json:"rating"`
}
