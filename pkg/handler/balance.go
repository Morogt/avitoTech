package handler

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func (h *Handler) info(c *gin.Context) {
	id, err := getUserById(c)
	if err != nil {
		return
	}

	info, err := h.services.Balance.GetInfo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      info.ID,
		"balance": info.Balance,
	})
}

func (h *Handler) addMoney(c *gin.Context) {
	id, count, err := getUserAndCount(c)
	if err != nil {
		return
	}
	count = math.Round(count*100) / 100

	info, err := h.services.Balance.AddMoney(id, count)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      info.ID,
		"balance": info.Balance,
		"added":   count,
	})
}

func (h *Handler) withdrawalOfMoney(c *gin.Context) {
	id, count, err := getUserAndCount(c)
	if err != nil {
		return
	}
	count = math.Round(count*100) / 100

	info, err := h.services.Balance.WithdrawalOfMoney(id, count)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       info.ID,
		"balance":  info.Balance,
		"withdraw": count,
	})
}

func (h *Handler) transfer(c *gin.Context) {
	senderId, recipientId, count, err := getSenderRecipientAndCount(c)
	if err != nil {
		return
	}
	count = math.Round(count*100) / 100

	senderLat, recipientLat, err := h.services.Balance.Transfer(senderId, recipientId, count)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"senderId":         senderLat.ID,
		"senderBalance":    senderLat.Balance,
		"recipientId":      recipientLat.ID,
		"recipientBalance": recipientLat.Balance,
	})
}

func (h *Handler) pay(c *gin.Context) {
	userId, serviceId, count, description, err := getUserServiceDescAndCount(c)
	if err != nil {
		return
	}
	count = math.Round(count*100) / 100

	user, orderId, err := h.services.Balance.Pay(userId, serviceId, count, description)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"orderId": orderId,
		"userId":  user.ID,
		"balance": user.Balance,
	})
}

func (h *Handler) history(c *gin.Context) {
	id, err := getUserById(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	history, err := h.services.Balance.History(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"": history,
	})
}
