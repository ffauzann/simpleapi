package request

type Login struct {
	Username string `json:"username" example:"admin1"`
	Password string `json:"password" example:"admin1"`
}
