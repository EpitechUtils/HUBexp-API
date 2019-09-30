package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	router := iris.New()
	router.Use(logger.New())
	router.Use(recover.New())
	router.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{"hello": "sailor"})
	})

	router.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
 