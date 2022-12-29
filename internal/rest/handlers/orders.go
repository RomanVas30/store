package handlers

import (
	"fmt"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/rest/middlewares"
	"github.com/RomanVas30/store/internal/rest/response"
	"github.com/RomanVas30/store/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewOrder(s service.Orders) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := middlewares.GetUserId(c)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var orderName struct {
			Name string `json:"name"`
		}

		if err := c.BindJSON(&orderName); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid order name: %v", err))
			return
		}

		order := entities.Order{Name: orderName.Name, UserId: userId}

		if err := s.CreateOrder(&order); err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": order.Id,
		})
	}
}

func GetOrders(s service.Orders) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := middlewares.GetUserId(c)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		orders, err := s.GetOrders(userId)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, *orders)
	}
}

func GetOrderById(s service.Orders) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := middlewares.GetUserId(c)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}

		order, err := s.GetOrderById(id, userId)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, *order)
	}
}

func OrderPayment(s service.Orders) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := middlewares.GetUserId(c)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}

		err = s.OrderPayment(id, userId)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]string{"status": "Success"})
	}
}

func AddProductToOrder(s service.Orders) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := middlewares.GetUserId(c)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var addProduct entities.AddProduct

		if err := c.BindJSON(&addProduct); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid parameters: %v", err))
			return
		}

		addProduct.UserId = userId

		err = s.AddProductToOrder(&addProduct)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]string{"status": "Success"})
	}
}
