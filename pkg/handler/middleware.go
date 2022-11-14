package handler

import (
	"avitoTech"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserById(c *gin.Context) (int, error) {
	var input avitoTech.Id
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	return int(input.ID), nil
}

func getServiceById(c *gin.Context) (int, error) {
	var input avitoTech.Id
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "service id not found")
		return 0, errors.New("service id not found")
	}
	return int(input.ID), nil
}

func getUserAndCount(c *gin.Context) (int, float64, error) {
	var input avitoTech.UserAndCount
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, 0, errors.New("user id not found")
	}
	return int(input.ID), input.Count, nil
}

func getSenderRecipientAndCount(c *gin.Context) (int, int, float64, error) {
	var input avitoTech.SenderRecipientAndCount
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, 0, 0, errors.New("sender id or recipient id not found")
	}
	return int(input.Sender), int(input.Recipient), input.Count, nil
}

func getUserServiceDescAndCount(c *gin.Context) (int, int, float64, string, error) {
	var input avitoTech.UserServiceAndCount
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, 0, 0, "", errors.New("user id or service id not found")
	}
	return int(input.UserId), int(input.ServiceId), input.Count, input.Description, nil
}
