package main

import (
	"net/http"
	"strconv"

	"github.com/EnggarSe/http-service/model"
	"github.com/labstack/echo"
)

func main() {
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
		id, _ := strconv.Atoi(c.Param("id"))
		article := store.ArticleMap
		return c.JSON(http.StatusOK, article[id-1])
	})
	e.DELETE("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		article := store.Remove(id)
		return c.JSON(http.StatusOK, article)
	})
	e.PUT("/articles/:id", func(c echo.Context) error {
		title := c.FormValue("Title")
		body := c.FormValue("Body")
		id, _ := strconv.Atoi(c.Param("id"))
		article := store.EditArticle(title, body, id)
		return c.JSON(http.StatusOK, article)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
