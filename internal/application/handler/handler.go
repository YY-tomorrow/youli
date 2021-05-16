package handler

import "anest.top/youli/internal/domain/external"

type Handler struct {
	register *external.Register
}

func NewHandler(r *external.Register) *Handler {
	return &Handler{r}
}
