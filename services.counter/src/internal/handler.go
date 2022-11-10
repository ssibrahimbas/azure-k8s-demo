package internal

import (
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/validator"
)

type Handler struct {
	s *Srv
	http *http.Client
	v *validator.Validator
	i18n *i18n.I18n
}

type HandlerParams struct {
	Srv *Srv
	Http  *http.Client
	I18n  *i18n.I18n
	Valid *validator.Validator
}

func NewHandler(params *HandlerParams) *Handler {
	return &Handler{
		s: params.Srv,
		http: params.Http,
		v: params.Valid,
		i18n: params.I18n,
	}
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func(h *Handler) initV1() {
	v1 := h.http.App.Group("api/counter/v1")
	v1.Get("/", h.Get)
	v1.Post("/inc", h.Inc)
	v1.Post("/dec", h.Dec)
	v1.Post("/reset", h.Reset)
	v1.Post("/set", h.Set)
}