package main

import (
	"github.com/gin-gonic/gin"
	"github.com/silvajuniorow/microAPI/cmd/api/controllers"
	"github.com/silvajuniorow/microAPI/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/api/v1/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategories(ctx, inMemoryCategoryRepository)
	})

}
