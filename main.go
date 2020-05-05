package main

import (
	"net/http"
	"strconv"

	"github.com/EnggarSe/http-service/model"
	"github.com/labstack/echo"
)

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	user := c.QueryParam("user") // bisa juga menggunakan FormValue namun form value digunakan untuk put patch dan post
	// 	age := c.QueryParam("age")
	// 	return c.String(http.StatusOK, user+","+age)
	// })
	// e.Logger.Fatal(e.Start(":8080"))

	// e := echo.New()
	// e.GET("articles", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Get List")
	// })
	// e.POST("articles", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Create List")
	// })
	// e.PUT("articles/:id", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Update List")
	// })
	// e.DELETE("articles/:id", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Delete List")
	// })
	// e.GET("articles/:id", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Get List by id")
	// })

	store := model.NewArticleStoreInMemory()
	e := echo.New()

	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap
		return c.JSON(http.StatusOK, articles)
	})
	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("Title")
		body := c.FormValue("Body")
		article, _ := model.CreateArticle(title, body)
		store.Save(article)
		return c.JSON(http.StatusOK, article)
	})
	e.GET("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		idParse, _ := strconv.Atoi(id)
		article := store.ArticleMap
		return c.JSON(http.StatusOK, article[idParse])
	})
	e.DELETE("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		idParse, _ := strconv.Atoi(id)
		article := store.Remove(idParse)
		return c.JSON(http.StatusOK, article)
	})
	e.PUT("/articles/:id", func(c echo.Context) error {
		title := c.FormValue("Title")
		body := c.FormValue("Body")
		id := c.Param("id")
		idParse, _ := strconv.Atoi(id)
		article := store.EditArticle(title, body, idParse)
		return c.JSON(http.StatusOK, article)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
