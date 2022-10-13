package repository

import (
	"database/sql"

	"github.com/vninomtz/worktion/pkg/model"
)

type sqliteExercise struct {
  DB *sql.DB
}

func NewSqliteExerciseRepo(db *sql.DB) model.ExerciseRepository {
  return &sqliteExercise {
    DB: db,
  }
}

func (r *sqliteExercise) FindByName(query string) ([]*model.Exercise, error) {

  return nil, nil
}
