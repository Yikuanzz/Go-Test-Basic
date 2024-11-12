package handler

import (
	"go-test-basic/common"
	"go-test-basic/model"
	"go-test-basic/service"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateResponse struct {
	ID int `json:"id"`
}

func CreateItem(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	item := &model.Item{
		Name:        req.Name,
		Description: req.Description,
	}

	err := service.CreateItem(c.Request.Context(), item)
	if err != nil {
		c.JSON(common.StatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, CreateResponse{
		ID: item.ID,
	})
}
