package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stassme/GoTask1/internal/messagesService"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *messagesService.MessageService
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessagesHandler(c echo.Context) error {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, messages)
}

func (h *Handler) PostMessageHandler(c echo.Context) error {
	var message messagesService.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, createdMessage)
}

func (h *Handler) PatchMessageHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	var updatedMessage messagesService.Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to decode JSON: " + err.Error()})
	}

	message, err := h.Service.UpdateMessageByID(id, updatedMessage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update message: " + err.Error()})
	}

	return c.JSON(http.StatusOK, message)
}

func (h *Handler) DeleteMessageHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	if err := h.Service.DeleteMessageByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
