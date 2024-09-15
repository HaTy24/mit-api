package request

type LoginRequest struct {
	Email    string `json:"email" example:"hathienty1@gmail.com"`
	Password string `json:"password" example:"123456"`
}

type SignUpRequest struct {
	FirstName   string `json:"first_name" example:"Ty"`
	LastName    string `json:"last_name" example:"Ha"`
	Email       string `json:"email" example:"hathienty1@gmail.com"`
	PhoneNumber string `json:"phone_number" example:"0948162501"`
	Password    string `json:"password" example:"123456"`
}
