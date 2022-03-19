package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	// 0 => User
	// 1 => Admin
	// 2 => Super Admin

	fmt.Println("Rest API Echo")

	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	//e.GET("/path", functionName)
	//e.GET("/path/:idParam", functionName)

	e.GET("/users/pw", GetAllUser)
	e.GET("/users", GetUserName)

	e.POST("/users", InsertUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)

	e.Logger.Fatal(e.Start(":425"))
}

/*
	Send Paramter
	*Path Param
		/path/idParam
		e.GET("/path/:idParam", functionName)
		id := c.Param("idParam")
	*Query Param
		/path?namaParam=Adithia
		e.GET("/path", functionName)
		namaUser := c.QueryParam("namaParam")
	*Form application/x-www-form-urlencoded
		namaParam => Adithia
		e.POST("/path", functionName)
		namaUser := c.FormValue("namaParam")

*/
