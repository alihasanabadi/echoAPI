package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"echoAPI/models"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Logging(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := f(c); err != nil {
			c.Error(err)
		}

		myUrl := c.Request()
		fmt.Println(myUrl)
		return nil

	}

}

func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		/*
			if err := next(c); err != nil {
				c.Error(err)
			}
		*/
		//header := c.Request().Header
		//token := header.Get("Token")

		cookie, err := c.Cookie("username")
		if err != nil {
			return err
		}
		fmt.Println(cookie.Value)

		token, err := strconv.Atoi(cookie.Value)
		checkErr(err)
		db := models.SetupDB()
		rows, err := db.Query("SELECT   token FROM users WHERE token = $1;", token)
		checkErr(err)

		//token := c.Header("Token")
		//myUrl := c.Request()
		if !rows.Next() {

			fmt.Println("No Auth")
			fmt.Println(token)
			//return c.JSON(http.StatusUnauthorized, "No Auth")
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

		}
		fmt.Println(token)

		return next(c)
		//return c.JSON(http.StatusUnauthorized, "No Auth")

		//return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}

}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
