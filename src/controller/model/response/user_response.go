package response

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"Eemail"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
