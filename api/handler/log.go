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

// @Summary      Log Data Info
// @Description  Use log id get log data info
// @Tags         log
// @Param        log_id    path     string  true  "search by log id"
// @Produce      json
// @Success      200
// @Router       /api/v1/log/:log_id [get]
// @Success 200 {object} model.Log
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

// @Summary      List Log Data
// @Description  Get Log Info List from database
// @Tags         log
// @Param        pagination    query     model.Pagination  false  "name search by q"  Format(email)
// @Produce      json
// @Success      200
// @Router       /api/v1/log [get]
// @Success 200 {array} model.Log
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
