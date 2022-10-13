package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchExercise(c *gin.Context) {
	res, _ := h.ExerciseService.Search("")

	c.JSON(http.StatusOK, gin.H{"response": res})
}
