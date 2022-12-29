package handlers

import (
	"fmt"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/rest/response"
	"github.com/RomanVas30/store/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewProduct(s service.Products) func(c *gin.Context) {
	return func(c *gin.Context) {
		var product entities.Product

		if err := c.BindJSON(&product); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid product parameters: %v", err))
			return
		}

		if err := s.CreateProduct(&product); err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": product.Id,
		})
	}
}

func GetProducts(s service.Products) func(c *gin.Context) {
	return func(c *gin.Context) {
		products, err := s.GetProducts()
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, *products)
	}
}

func GetProductById(s service.Products) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}

		product, err := s.GetProductById(id)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, *product)
	}
}
