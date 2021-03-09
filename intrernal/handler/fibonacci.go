package handler

import (
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/intrernal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getFibonacci(c *gin.Context) {
	var input model.Fibonacci

	if err := c.Bind(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	ok := validateInputFibonacci(input, c)
	if !ok {
		return
	}

	result, err := h.services.Fibonacci.GetSlice(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func validateInputFibonacci(input model.Fibonacci, c *gin.Context) bool {
	ok := true
	if input.Y <= input.X {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Y should have more than X")
		ok = false
	} else if input.X > 92 || input.Y > 92 {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Y and X must be less than 92")
		ok = false
	}

	return ok
}
