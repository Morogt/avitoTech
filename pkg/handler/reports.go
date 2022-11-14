package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) get(c *gin.Context) {
	info, err := h.services.Services.GetReportByServ()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"": info,
	})
}
