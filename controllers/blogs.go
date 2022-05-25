package controllers

import (
	"net/http"
	"time"

	"echoAPI/models"

	"github.com/labstack/echo/v4"

	"strconv"
)

type CreatePostInput struct {
	Subject string `json:"subject" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type UpdatePostInput struct {
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetPost(c echo.Context) error {

	db := models.SetupDB()

	rows, err := db.Query("SELECT * FROM blogs")
	checkErr(err)

	var posts []models.Blog

	for rows.Next() {
		var id uint
		var subject string
		var content string

		err = rows.Scan(&id, &subject, &content)

		// check errors
		checkErr(err)

		posts = append(posts, models.Blog{ID: id, Subject: subject, Content: content})
	}

	return c.JSON(http.StatusOK, posts)
}

func UpdatePost(c echo.Context) (err error) {

	input := new(UpdatePostInput)
	if err = c.Bind(input); err != nil {
		return c.JSON(http.StatusOK, "Error Type")
	}

	// Get model if exist

	db := models.SetupDB()

	idinput := c.Param("id")

	rows, err := db.Query("SELECT id FROM blogs WHERE id = $1;", idinput)
	checkErr(err)

	if !rows.Next() {
		return c.JSON(http.StatusOK, "The Post Not nfound")
	}

	rowss, err := db.Exec("UPDATE blogs SET subject = $2, content = $3 where (id = $1);", idinput, input.Subject, input.Content)
	checkErr(err)

	return c.JSON(http.StatusOK, rowss)
}

func LoginUser(c echo.Context) (err error) {
	var Output int
	header := c.Request().Header
	email := header.Get("Email")
	password := header.Get("Password")
	db := models.SetupDB()
	rows, err := db.Query("SELECT password, token FROM users WHERE email = $1;", email)
	checkErr(err)

	//token := c.Header("Token")
	//myUrl := c.Request()
	if !rows.Next() {
		return echo.NewHTTPError(http.StatusUnauthorized, "The Email Or Password Is Wrong !!!")
	} else {

		var passwd string
		var token int

		err = rows.Scan(&passwd, &token)

		checkErr(err)

		if passwd == password {
			Output = token

		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "The Email Or Password Is Wrong !!!")
		}

	}
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = strconv.Itoa(Output)

	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, Output)
}
