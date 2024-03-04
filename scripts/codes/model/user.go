package model

type UserLoginResponse struct {
	Id         int    `json:"id"            ` //
	Username   string `json:"name"`
	Permission string `json:"permission"`
}
