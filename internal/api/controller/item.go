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

type ItemController struct {
	Service service.ItemService
}

func NewItemController(s service.ItemService) *ItemController {
	return &ItemController{
		Service: s,
	}
}

func (h *ItemController) Get(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = dto.GetItemRequest{}
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
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data Item", data, meta)
}

func (s *ItemController) GetById(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	data, err := s.Service.GetById(ctx, id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data item", data, nil)
}

func (h *ItemController) Create(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.CreateItemRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	err := h.Service.Create(ctx, request.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusCreated, "Berhasil menyimpan data item", nil, nil)
}

func (h *ItemController) Update(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		id, _   = uuid.Parse(c.Param("id"))
		request dto.UpdateItemRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	err := h.Service.Update(ctx, request.ParseToModel(), id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil memperbarui data item", nil, nil)
}

func (h *ItemController) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	if err := h.Service.Delete(ctx, id); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil menghapus data item", nil, nil)
}
