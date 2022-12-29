package middlewares

import (
	"fmt"
	"github.com/RomanVas30/store/internal/rest/response"
	"github.com/RomanVas30/store/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func UserIdentity(s service.Authorization, isAdmin bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			response.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
			return
		}

		userId, userRole, err := s.ParseToken(headerParts[1])
		if err != nil {
			response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if isAdmin && userRole != "admin" {
			response.NewErrorResponse(c, http.StatusUnauthorized, "insufficient permissions to perform this operation")
			return
		}

		c.Set(userCtx, userId)
	}
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, fmt.Errorf("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, fmt.Errorf("user id is of invalid type")
	}

	return idInt, nil
}
