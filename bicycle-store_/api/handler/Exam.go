package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ViewDate(c *gin.Context) {
	date := c.Param("time")

	fmt.Println(date)

	resp, err := h.storages.Code().GetDate(context.Background(), &models.GiveMe{Day: date})
	if err != nil {
		h.handlerResponse(c, "storage.code.getDate", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get otchet", http.StatusCreated, resp)

}
