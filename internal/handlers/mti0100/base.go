package mti0100

import "gatekeeper-core/internal/transport"

type Handler struct {
	client transport.Client
}

func New(client transport.Client) *Handler {
	return &Handler{client: client}
}
