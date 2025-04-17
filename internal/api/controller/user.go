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

type UserController struct {
	Service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{
		Service: s,
	}
}

func (h *UserController) Get(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = dto.GetUserRequest{}
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
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data user", data, meta)
}

func (h *UserController) GetById(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	data, err := h.Service.GetById(ctx, id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data user", data, nil)
}

func (h *UserController) Create(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.UserRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	data, err := h.Service.Create(ctx, request.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusCreated, "Berhasil menyimpan data user", map[string]string{"id": data}, nil)
}

func (h *UserController) Update(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		id, _   = uuid.Parse(c.Param("id"))
		request dto.UserRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	data, err := h.Service.Update(ctx, request.ParseToModel(), id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil memperbarui data user", map[string]string{"id": data}, nil)
}

func (h *UserController) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	if err := h.Service.Delete(ctx, id); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil menghapus data user", nil, nil)
}
