package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/fxivan/routeOptimizer/cockroach/models"
	"github.com/fxivan/routeOptimizer/cockroach/usecases"
)

type cockroachHttpHandler struct {
	cockroachUsecase usecases.CockroachUsecase
}

func NewCockroachHttpHandler(cockroachUsecase usecases.CockroachUsecase) CockroachHandler {
	return &cockroachHttpHandler{
		cockroachUsecase: cockroachUsecase,
	}
}

func (h *cockroachHttpHandler) DetectCockroach(c echo.Context) error {
	reqBody := new(models.AddCockroachData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	if err := h.cockroachUsecase.CockroachDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing data failed")
	}

	return response(c, http.StatusOK, "Success 🪳🪳🪳")
}

func (h *cockroachHttpHandler) TransactionSearchCockroach(c echo.Context) error {
	reqBody := new(models.IdCockroachData)
	fmt.Print("Req body -->", reqBody)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	resultResponse, err := h.cockroachUsecase.CockroachSearchTransactions(reqBody)

	if err != nil {
		return response(c, http.StatusInternalServerError, "Processing data failed")
	}

	responseMessage := fmt.Sprintf("%v", resultResponse)

	return response(c, http.StatusOK, responseMessage)
}
