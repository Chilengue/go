package request

type UserRequest struct {
	Email    string `json:"Eemail"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}
