package handlers

import (
	"fmt"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/rest/response"
	"github.com/RomanVas30/store/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewOrgUnit(s service.OrgUnits) func(c *gin.Context) {
	return func(c *gin.Context) {
		var unit entities.OrgUnit

		if err := c.BindJSON(&unit); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid staffer parameters: %v", err))
			return
		}

		if err := s.CreateOrgUnit(&unit); err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": unit.Id,
		})
	}
}

func GetOrgUnits(s service.OrgUnits) func(c *gin.Context) {
	return func(c *gin.Context) {
		units, err := s.GetOrgUnits()
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, *units)
	}
}

func DeleteOrgUnit(s service.OrgUnits) func(c *gin.Context) {
	return func(c *gin.Context) {

		var unit entities.OrgUnit

		if err := c.BindJSON(&unit); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid parameters: %v", err))
			return
		}

		if unit.Name == "" {
			response.NewErrorResponse(c, http.StatusBadRequest, "unit name should not be empty")
			return
		}

		err := s.DeleteOrgUnit(unit.Name)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{})
	}
}

func UpdateOrgUnit(s service.OrgUnits) func(c *gin.Context) {
	return func(c *gin.Context) {
		var unit entities.OrgUnit

		if err := c.BindJSON(&unit); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid staffer parameters: %v", err))
			return
		}

		if err := s.UpdateOrgUnit(&unit); err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": unit.Id,
		})
	}
}
