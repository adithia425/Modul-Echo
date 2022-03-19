package controllers

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo"
// )

// func SendMessageError(w http.ResponseWriter, message string) {
// 	var response ErrorResponse
// 	response.Status = 400
// 	response.Message = message

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func SendMessageSuccess(w http.ResponseWriter, message string) {
// 	var response ErrorResponse
// 	response.Status = 200
// 	response.Message = message

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func GetAllUser(c echo.Context) error {
// 	db := Connect()
// 	defer db.Close()

// 	query := "SELECT * FROM users"
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	var user User
// 	var users []User
// 	for rows.Next() {
// 		if err != rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password) {
// 			// error response
// 			log.Println(err)
// 			return err
// 		} else {
// 			users = append(users, user)
// 		}
// 	}

// 	// success response
// 	var response UserResponse
// 	response.Status = 200
// 	response.Message = "Request success"
// 	response.Users = users

// 	db.Close()
// 	if err := c.Bind(response); err != nil {
// 		return err
// 	}
// 	//return c.JSON(http.StatusOK, response)

// 	// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
// 	// c.Response().WriteHeader(http.StatusOK)
// 	// return json.NewEncoder(c.Response()).Encode(response)

// 	return c.JSONPretty(http.StatusOK, response, "  ")

//}
