package utils
type User struct {
	Id 	int `json:"id"`
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}
type Response struct {
	Status 		string `json:"status"`
	Message 	string `json:"message"`
	Data  User `json:"data"`
}
