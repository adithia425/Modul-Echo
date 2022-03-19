package main

import (
	"echo/controllers"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ============================================================================= Response

func SendMessageError(c echo.Context, message string) error {

	var response controllers.ErrorResponse
	response.Status = 400
	response.Message = message

	return c.JSONPretty(http.StatusBadRequest, response, "  ")
}
func SendMessageSuccess(c echo.Context, message string) error {

	var response controllers.SuccessResponse
	response.Status = 200
	response.Message = message

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func SendUserResponse(c echo.Context, user controllers.User) error {

	var response controllers.UserResponse
	response.Status = 200
	response.Message = "Request Success"
	response.User = user

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func SendAllUserResponse(c echo.Context, users []controllers.User) error {

	var response controllers.AllUserResponse
	response.Status = 200
	response.Message = "Request Success"
	response.Users = users

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func SendAllUserResponseNoPassword(c echo.Context, users []controllers.UserNoPassword) error {

	var response controllers.AllUserResponseNoPassword
	response.Status = 200
	response.Message = "Request Success"
	response.Users = users

	return c.JSONPretty(http.StatusOK, response, "  ")
}

// ============================================================================= GET
func GetAllUser(c echo.Context) error {
	db := controllers.Connect()
	defer db.Close()

	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return err
	}

	var user controllers.User
	var users []controllers.User
	for rows.Next() {
		if err != rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password) {
			// error response
			log.Println(err)
			return err
		} else {
			users = append(users, user)
		}
	}

	return SendAllUserResponse(c, users)
}

func GetUserName(c echo.Context) error {
	db := controllers.Connect()
	defer db.Close()

	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return err
	}

	var userTemp controllers.User
	var user controllers.UserNoPassword
	var users []controllers.UserNoPassword
	for rows.Next() {
		if err != rows.Scan(&userTemp.Id, &userTemp.Name, &userTemp.Email, &userTemp.Password) {
			// error response
			log.Println(err)
			return err
		} else {
			user.Id = userTemp.Id
			user.Name = userTemp.Name
			user.Email = userTemp.Email
			users = append(users, user)
		}
	}

	db.Close()
	return SendAllUserResponseNoPassword(c, users)
}

// ============================================================================= POST
func InsertUser(c echo.Context) error {
	db := controllers.Connect()
	defer db.Close()

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	_, errQuery := db.Exec("INSERT INTO users(name,email,password) VALUES (?,?,?)",
		name,
		email,
		password,
	)

	if errQuery != nil {
		log.Println(errQuery)
		return SendMessageError(c, "Query Error")
	}

	return SendMessageSuccess(c, "Success Insert New User")
}

// ============================================================================= PUT
func UpdateUser(c echo.Context) error {
	db := controllers.Connect()
	defer db.Close()

	id := c.Param("id")

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var Query string
	Query = "UPDATE users SET "

	if name != "" {
		Query = Query + "name = '" + name + "', "
	}
	if email != "" {
		Query = Query + "email = '" + email + "', "
	}
	if password != "" {
		Query = Query + "password = '" + password + "' "
	}

	Query = Query + "WHERE id = " + id

	_, errQuery := db.Exec(Query)
	if errQuery != nil {
		return SendMessageError(c, "Query Error")
	}

	var user controllers.User
	user.Id, _ = strconv.Atoi(id)
	user.Name = name
	user.Email = email
	user.Password = password

	return SendUserResponse(c, user)
}

// ============================================================================= DELETE
func DeleteUser(c echo.Context) error {
	db := controllers.Connect()
	defer db.Close()

	id := c.Param("id")

	Query := "DELETE FROM users WHERE id = " + id

	_, errQuery := db.Exec(Query)
	if errQuery != nil {
		return SendMessageError(c, "Query Error")
	}

	return SendMessageSuccess(c, "Delete User Success")
}

// ============================================================================= COOKIE
// func WriteCookie(c echo.Context) error {
// 	cookie := new(http.Cookie)
// 	cookie.Name = "username"
// 	cookie.Value = "jon"
// 	cookie.Expires = time.Now().Add(24 * time.Hour)
// 	c.SetCookie(cookie)
// 	return c.String(http.StatusOK, "write a cookie")
// }
// func readCookie(c echo.Context) error {
// 	cookie, err := c.Cookie("username")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(cookie.Name)
// 	fmt.Println(cookie.Value)
// 	return c.String(http.StatusOK, "read a cookie")
// }
