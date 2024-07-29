package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/tenhou/api/service"
	"github.com/gsxhnd/tenhou/model"
	"github.com/gsxhnd/tenhou/utils"
)

type LogHandler interface {
	GetLogInfoByLogId(ctx *fiber.Ctx) error
	GetLogInfoList(ctx *fiber.Ctx) error
}

type logHandler struct {
	validator *validator.Validate
	svc       service.LogService
	logger    utils.Logger
}

func NewLogHandler(v *validator.Validate, l utils.Logger, svc service.LogService) LogHandler {
	return &logHandler{
		svc:       svc,
		validator: v,
		logger:    l,
	}
}

// @Description  Get Log Info from database by log id
// @Produce      json
// @Success      200
// @Router       /log/:log_id [get]
func (h *logHandler) GetLogInfoByLogId(ctx *fiber.Ctx) error {
	var logId = ctx.Params("log_id")
	err := h.validator.Var(logId, "gt=10")
	if err != nil {
		h.logger.Errorw("url validator error", "error", err.Error())
		return ctx.Status(500).SendString(err.Error())
	}

	data, err := h.svc.GetLogInfoByLogId(logId)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

// @Description  Get Log Info List from database
// @Produce      json
// @Success      200
// @Router       /log [get]
func (h *logHandler) GetLogInfoList(ctx *fiber.Ctx) error {
	var p = model.Pagination{}
	if err := ctx.QueryParser(&p); err != nil {
		h.logger.Errorw("url validator error", "error", err.Error())
		return ctx.Status(500).SendString(err.Error())
	}

	h.logger.Debugf("url query: %#v", p)

	data, err := h.svc.GetLogInfoList(p)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}
