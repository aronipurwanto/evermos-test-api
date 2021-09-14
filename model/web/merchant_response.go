package web

type MerchantResponse struct {
	Id   		int    `json:"id"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Address 	string `json:"address"`
	Rating	 	float64 `json:"rating"`
}
