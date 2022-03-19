package controllers

type User struct {
	Id       int    `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type UserNoPassword struct {
	Id    int    `json:"id" form:"id" query:"id"`
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	User    User   `json:"users"`
}

type UserResponseNoPassword struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	User    UserNoPassword `json:"users"`
}

type AllUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Users   []User `json:"users"`
}

type AllUserResponseNoPassword struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Users   []UserNoPassword `json:"users"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
