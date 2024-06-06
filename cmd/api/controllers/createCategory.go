package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silvajuniorow/microAPI/internal/repositories"
	"github.com/silvajuniorow/microAPI/internal/usecases"
)

type createCategoryInputDTO struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInputDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase(repository)

	err := useCase.Execute(body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
