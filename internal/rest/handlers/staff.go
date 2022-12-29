package handlers

import (
	"fmt"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/rest/response"
	"github.com/RomanVas30/store/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewStaffer(s service.Staff) func(c *gin.Context) {
	return func(c *gin.Context) {
		var staffer entities.Staffer

		if err := c.BindJSON(&staffer); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid staffer parameters: %v", err))
			return
		}

		if err := s.CreateStaffer(&staffer); err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id":              staffer.Id,
			"employment_date": staffer.EmploymentDate,
		})
	}
}

func GetStaff(s service.Staff) func(c *gin.Context) {
	return func(c *gin.Context) {
		staff, err := s.GetStaff()
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, *staff)
	}
}

func DeleteStaffer(s service.Staff) func(c *gin.Context) {
	return func(c *gin.Context) {

		var staffer entities.ActionStaffer

		if err := c.BindJSON(&staffer); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid parameters: %v", err))
			return
		}

		if staffer.SNILS == "" {
			response.NewErrorResponse(c, http.StatusBadRequest, "staffer snils should not be empty")
			return
		}

		err := s.DeleteStaffer(staffer.SNILS)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{})
	}
}

func SearchStaff(s service.Staff) func(c *gin.Context) {
	return func(c *gin.Context) {

		var searchInput entities.ActionStaffer

		if err := c.BindJSON(&searchInput); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid parameters: %v", err))
			return
		}

		if searchInput.FIO == "" && searchInput.SNILS == "" {
			response.NewErrorResponse(c, http.StatusBadRequest, "fio or snils should not be empty")
			return
		}

		staff, err := s.SearchStaff(searchInput.FIO, searchInput.SNILS)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, *staff)
	}
}

func UpdateStaffer(s service.Staff) func(c *gin.Context) {
	return func(c *gin.Context) {

		var updateStaffer entities.ActionStaffer

		if err := c.BindJSON(&updateStaffer); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest,
				fmt.Sprintf("invalid parameters: %v", err))
			return
		}

		staffer := entities.NewStafferByUpdateStaffer(&updateStaffer)

		err := s.UpdateStaffer(staffer, updateStaffer.AddPosts, updateStaffer.DeletePosts)
		if err != nil {
			response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{})
	}
}
