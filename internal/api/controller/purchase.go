package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/dto"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	service "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type PurchaseController struct {
	Service service.PurchaseService
}

func NewPurchaseController(s service.PurchaseService) *PurchaseController {
	return &PurchaseController{
		Service: s,
	}
}

// @Summary Get purchases
// @Description Mendapatkan semua data pembelian
// @Tags purchases
// @Accept json
// @Produce json
// @Param data body dto.GetPurchaseRequest true "Purchase request"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Router /purchases [get]
func (h *PurchaseController) Get(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = dto.GetPurchaseRequest{}
	)

	if err := helper.BindRequest(c, &payload, true); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	if payload.Limit == 0 {
		payload.Limit = 10
	}
	if payload.Page == 0 {
		payload.Page = 1
	}

	// data, err := h.Service.Get(ctx)
	data, total, err := h.Service.Get(ctx, payload.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	meta := helper.NewMetadata(payload.Page, payload.Limit, len(data), int(total))
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data pembelian", data, meta)
}

// @Summary Create bulk purchase
// @Description Buat data pembelian secara bulk maupun satu per satu
// @Tags purchases
// @Accept json
// @Produce json
// @Param data body dto.BulkPurchaseRequest true "Purchase request"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Router /purchases [post]
func (h *PurchaseController) CreateBulk(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.BulkPurchaseRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	for _, item := range request.Data {
		if item.ItemID == uuid.Nil {
			return response.BuildErrorResponse(c, response.NewCustomError(http.StatusBadRequest, "Item ID tidak boleh kosong", nil))
		}

		if item.PurchasedQty < 1 {
			return response.BuildErrorResponse(c, response.NewCustomError(http.StatusBadRequest, "Jumlah pembelian tidak boleh kurang dari 1", nil))
		}
	}

	requests := make([]model.StockBatchItem, len(request.Data))
	for i, item := range request.Data {
		requests[i] = item.ParseToModel()
	}

	if err := h.Service.CreateBulk(ctx, requests); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusCreated, "Berhasil menyimpan data pembelian", nil, nil)
}

// @Summary Update purchases
// @Description Memperbarui data pembelian
// @Tags purchases
// @Accept json
// @Produce json
// @Param data body dto.StockBatchItem true "Purchase request"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Router /purchases/:id [put]
func (h *PurchaseController) Update(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		id, _   = uuid.Parse(c.Param("id"))
		request dto.StockBatchItem
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	if request.ItemID == uuid.Nil {
		return response.BuildErrorResponse(c, response.NewCustomError(http.StatusBadRequest, "Item ID tidak boleh kosong", nil))
	}

	parsedRequest := request.ParseToModel()
	err := h.Service.Update(ctx, &parsedRequest, id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil memperbarui data pembelian", nil, nil)
}

// @Summary Delete purchases
// @Description Menghapus data pembelian berdasarkan id
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path string true "ID Pembelian"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Router /purchases/:id [delete]
func (h *PurchaseController) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	if err := h.Service.Delete(ctx, id); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil menghapus data pembelian", nil, nil)
}
