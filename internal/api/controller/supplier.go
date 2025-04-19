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

type SupplierController struct {
	Service service.SupplierService
}

func NewSupplierController(s service.SupplierService) *SupplierController {
	return &SupplierController{
		Service: s,
	}
}

func (h *SupplierController) Get(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = dto.GetSupplierRequest{}
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
	if payload.Sort == "" {
		payload.Sort = "name"
	}
	if payload.Order == "" {
		payload.Order = "asc"
	}

	data, total, err := h.Service.Get(ctx, payload.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	meta := helper.NewMetadata(payload.Page, payload.Limit, len(data), int(total))
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data supplier", data, meta)
}

func (s *SupplierController) GetById(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	data, err := s.Service.GetById(ctx, id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data supplier", data, nil)
}

func (h *SupplierController) Create(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.SupplierRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	err := h.Service.Create(ctx, request.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusCreated, "Berhasil menyimpan data supplier", nil, nil)
}

func (h *SupplierController) Update(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		id, _   = uuid.Parse(c.Param("id"))
		request dto.SupplierRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	err := h.Service.Update(ctx, request.ParseToModel(), id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil memperbarui data supplier", nil, nil)
}

func (h *SupplierController) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	if err := h.Service.Delete(ctx, id); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil menghapus data supplier", nil, nil)
}
