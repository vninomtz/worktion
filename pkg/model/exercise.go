package model

// Domanin model entity: Exercise

type Exercise struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Muscles string `json:"muscles"`
}
