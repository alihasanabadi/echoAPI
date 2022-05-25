package main

import (
	_ "net/http"

	"github.com/labstack/echo/v4"

	"echoAPI/controllers"

	"echoAPI/middlewares"
)

func main() {
	e := echo.New()

	//AuthGetPost := echo.HandlerFunc(controllers.GetPost)

	//e.Use(middlewares.BasicAuthWithConfig(middlewares.BasicAuthConfig{}))

	//e.Use(middlewares.BasicAuth(AuthGetPost))
	g := e.Group("/posts")

	g.Use(middlewares.BasicAuth)
	/*
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				// Extract the credentials from HTTP request header and perform a security
				// check

				// For invalid credentials
				return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

				// For valid credentials call next
				// return next(c)
			}
		})
	*/
	//e.GET("/posts", middlewares.BasicAuth(controllers.GetPost))
	g.GET("", controllers.GetPost)
	g.PUT("/:id", controllers.UpdatePost)

	e.POST("/login", controllers.LoginUser)
	//e.GET("/postsp", middlewares.ServerHeader(controllers.GetPost))
	e.Logger.Fatal(e.Start(":8080"))
}
