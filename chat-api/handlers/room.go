package handlers

import (
	"chatjobsity/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	roomService services.RoomService
}

func NewRoomHandler(roomService services.RoomService) *RoomHandler {
	return &RoomHandler{roomService: roomService}
}

func (h *RoomHandler) Rooms(c *gin.Context) {
	rooms, err := h.roomService.Rooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting rooms"})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *RoomHandler) Messages(c *gin.Context) {
	roomId := c.Param("roomId")

	messages, err := h.roomService.Messages(roomId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
