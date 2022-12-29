package handlers

import (
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/rest/response"
	"github.com/RomanVas30/store/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(s service.Authorization) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input entities.User

		if err := c.BindJSON(&input); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
			return
		}

		id, err := s.CreateUser(input)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func SignIn(s service.Authorization) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input entities.UserCred

		if err := c.BindJSON(&input); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		token, err := s.GenerateToken(input)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
		})
	}
}

func ChangePassword(s service.Authorization) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input entities.ChangePassword

		if err := c.BindJSON(&input); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		err := s.ChangePassword(input)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{})
	}
}
