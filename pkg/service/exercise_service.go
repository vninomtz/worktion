package service

import "github.com/vninomtz/worktion/pkg/model"

type ExerciseServiceImp struct {
  repo model.ExerciseRepository
}

func NewExerciseService(r model.ExerciseRepository) model.ExerciseService {
  return &ExerciseServiceImp{
    repo: r,
  }
}

func (s *ExerciseServiceImp) Search(query string) (*[]model.Exercise, error) {

  res := []model.Exercise{
    model.Exercise{ID: "1", Name: "Pull ups",Muscles: "Biceps"},
    model.Exercise{ID: "2", Name: "Chin ups",Muscles: "Biceps"},
    model.Exercise{ID: "3", Name: "Push ups",Muscles: "Biceps"},
    model.Exercise{ID: "4", Name: "Dips",Muscles: "Biceps"},
  }

  return &res, nil
}
