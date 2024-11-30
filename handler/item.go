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

	// To test
	// config := configService.GetConfig()
	// if len(req.Name) > config.MaxLength {
	// 	c.JSON(400, gin.H{"error": "Name is too long"})
	// 	return
	// }

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

type GetRequest struct {
	ID int `json:"id"`
}

type GetResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetItem(c *gin.Context) {
	var req GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	item, err := service.GetItem(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(common.StatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, GetResponse{
		Name:        item.Name,
		Description: item.Description,
	})
}

type ListResponse struct {
	Items []*GetResponse `json:"items"`
}

func ListItems(c *gin.Context) {
	items, err := service.ListItems(c.Request.Context())
	if err != nil {
		c.JSON(common.StatusCode(err), gin.H{"errors": err.Error()})
		return
	}

	var resp ListResponse
	for _, item := range items {
		resp.Items = append(resp.Items, &GetResponse{
			Name:        item.Name,
			Description: item.Description,
		})
	}

	c.JSON(200, resp)
}

type UpdateRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func UpdateItem(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	item := &model.Item{
		Name:        req.Name,
		Description: req.Description,
	}

	err := service.UpdateItem(c.Request.Context(), req.ID, item)
	if err != nil {
		c.JSON(common.StatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, nil)
}

type DeleteRequest struct {
	ID int `json:"id"`
}

func DeleteItem(c *gin.Context) {
	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := service.DeleteItem(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(common.StatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, nil)
}
