package handler

import (
	"avitoTech"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	var input avitoTech.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Users.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
