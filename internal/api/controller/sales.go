package controller

import (
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
