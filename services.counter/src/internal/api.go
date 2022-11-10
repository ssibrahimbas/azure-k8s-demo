package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/azure-k8s-demo.counter/src/dto"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	count := h.s.Get()
	l, a := h.i18n.GetLanguagesInContext(c)
	return result.SuccessData(h.i18n.Translate("success_counter_get", l, a), count, fiber.StatusOK)
}

func (h *Handler) Inc(c *fiber.Ctx) error {
	h.s.Inc()
	count := h.s.Get()
	l, a := h.i18n.GetLanguagesInContext(c)
	return result.SuccessData(h.i18n.Translate("success_counter_inc", l, a), count, fiber.StatusOK)
}

func (h *Handler) Dec(c *fiber.Ctx) error {
	h.s.Dec()
	count := h.s.Get()
	l, a := h.i18n.GetLanguagesInContext(c)
	return result.SuccessData(h.i18n.Translate("success_counter_dec", l, a), count, fiber.StatusOK)
}

func (h *Handler) Reset(c *fiber.Ctx) error {
	h.s.Reset()
	count := h.s.Get()
	l, a := h.i18n.GetLanguagesInContext(c)
	return result.SuccessData(h.i18n.Translate("success_counter_reset", l, a), count, fiber.StatusOK)
}

func (h *Handler) Set(c *fiber.Ctx) error {
	d := &dto.SetCount{}
	_ = c.BodyParser(d)
	l, a := h.i18n.GetLanguagesInContext(c)
	if err := c.BodyParser(d); err != nil {
		return result.Error(h.i18n.Translate("error_invalid_request", l, a), fiber.StatusBadRequest)
	}
	if errors := h.v.ValidateStruct(d, l,a); len(errors) > 0 {
		return result.ErrorData(h.i18n.Translate("error_invalid_validation", l, a), errors, fiber.StatusBadRequest)
	}
	h.s.Set(d.Count)
	count := h.s.Get()
	return result.SuccessData(h.i18n.Translate("success_counter_set", l, a), count, fiber.StatusOK)
}