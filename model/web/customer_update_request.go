package web

type CustomerUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
	Email 		string `validate:"required,min=1,max=100" json:"email"`
	Address 	string `validate:"required,min=1,max=100" json:"address"`
	PhoneNumber	string `validate:"required,min=1,max=100" json:"phone_number"`
}
