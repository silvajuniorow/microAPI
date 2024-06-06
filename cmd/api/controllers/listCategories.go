package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silvajuniorow/microAPI/internal/repositories"
	"github.com/silvajuniorow/microAPI/internal/usecases"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := usecases.NewListCategoriesUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"success": true, "categories": categories})
}
