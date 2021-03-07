package handler

import (
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/intrernal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getFibonacci(c *gin.Context)  {
	var input model.Fibonacci

	if err := c.Bind(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	result, err := h.services.Fibonacci.GetSlice(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
