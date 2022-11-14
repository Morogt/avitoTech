package handler

//func (h *Handler) addService(c *gin.Context) {
//	var input avitoTech.Service
//
//	if err := c.BindJSON(&input); err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	id, err := h.services.Services.CreateServ(input)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"id": id,
//	})
//}
