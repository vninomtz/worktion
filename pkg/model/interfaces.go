package model

type ExerciseService interface {
  Search(query string) (*[]Exercise, error)
}

type ExerciseRepository interface {
  FindByName(name string) ([]*Exercise, error)
}
