package handler

import (
	"github.com/RockkleyPushPost/api-gateway/core"
	"github.com/gofiber/fiber/v2"
)

type GatewayHandler struct {
	registry *core.ServiceRegistry
}

func NewGatewayHandler(registry *core.ServiceRegistry) *GatewayHandler {

	return &GatewayHandler{registry: registry}
}

func (h *GatewayHandler) HandleRequest(c *fiber.Ctx) error {

	return h.registry.ForwardRequest(c)
}
