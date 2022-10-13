package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vninomtz/worktion/pkg/model"
)

type Handler struct {
  ExerciseService model.ExerciseService
}

func NewHandler(r *gin.Engine, exerciseService model.ExerciseService){
  h := &Handler{
    ExerciseService: exerciseService,
  }

  g := r.Group("/api/v1")

  g.GET("exercise", h.SearchExercise)
}
