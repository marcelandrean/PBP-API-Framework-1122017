package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Age      string `json:"age" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserType int    `json:"usertype" validate:"required"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
