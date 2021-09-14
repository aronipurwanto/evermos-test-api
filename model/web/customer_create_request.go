package web

type CustomerCreateRequest struct {
	Name 		string `validate:"required,min=1,max=100" json:"name"`
	Email 		string `validate:"required,min=1,max=100" json:"email"`
	Address 	string `validate:"required,min=1,max=100" json:"address"`
	PhoneNumber	string `validate:"required,min=1,max=100" json:"phone_number"`
}
