package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/dto"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	service "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type SaleController struct {
	Service service.SaleService
}

func NewSaleController(s service.SaleService) *SaleController {
	return &SaleController{
		Service: s,
	}
}

// @Summary Get sales
// @Description Mendapatkan semua data penjualan
// @Tags sales
// @Accept json
// @Produce json
// @Param data body dto.GetSaleRequest true "Sale request"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Router /sales [get]
func (h *SaleController) Get(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = dto.GetSaleRequest{}
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

	data, total, err := h.Service.Get(ctx, payload.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	meta := helper.NewMetadata(payload.Page, payload.Limit, len(data), int(total))
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data penjualan", data, meta)
}

// @Summary Get sale by id
// @Description Mendapatkan data penjualan berdasarkan id
// @Tags sales
// @Accept json
// @Produce json
// @Param id path string true "ID Penjualan"
// @Success 200 {object} response.JSONResponse
// @Failure 404 {object} response.JSONResponse
// @Router /sales/:id [get]
func (s *SaleController) GetById(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	data, err := s.Service.GetById(ctx, id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data penjualan", data, nil)
}

// @Summary Create sale
// @Description Buat transaksi penjualan menggunakan metode FIFO
// @Tags sales
// @Accept json
// @Produce json
// @Param data body dto.SaleRequest true "Sale request"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Router /sales [post]
func (h *SaleController) Create(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.SaleRequest
	)

	userID, ok := c.Get("user_id").(string)
	if !ok || userID == "" {
		return response.BuildErrorResponse(c, response.NewCustomError(http.StatusUnauthorized, "User tidak terautentikasi", nil))
	}

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	if err := h.Service.Create(ctx, request.ParseToModel(userID)); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusCreated, "Berhasil menyimpan data penjualan", nil, nil)
}

// @Summary Delete sale
// @Description Menghapus data penjualan berdasarkan id
// @Tags sales
// @Accept json
// @Produce json
// @Param id path string true "ID Penjualan"
// @Success 200 {object} response.JSONResponse
// @Failure 404 {object} response.JSONResponse
// @Router /sales/:id [delete]
func (h *SaleController) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	if err := h.Service.Delete(ctx, id); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil menghapus data penjualan", nil, nil)
}

// @Summary Get monthly sales report
// @Description Mendapatkan data laporan total penjualan, total HPP, dan total profit
// @Tags sales
// @Accept json
// @Produce json
// @Param data body dto.GetMonthlySalesReport true "Sale request"
// @Success 200 {object} response.JSONResponse
// @Failure 400 {object} response.JSONResponse
// @Router /sales/report [get]
func (h *SaleController) GetMonthlySalesReport(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = dto.GetMonthlySalesReport{}
	)

	if err := helper.BindRequest(c, &payload, true); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	fmt.Println("Month:", payload.Month)
	fmt.Println("Year:", payload.Year)

	data, err := h.Service.GetMonthlySalesReport(ctx, payload.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data penjualan", data, nil)
}
