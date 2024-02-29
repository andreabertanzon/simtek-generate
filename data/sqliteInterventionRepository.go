package data

import "github.com/andreabertanzon/simtek-generate/models"

type SqliteInterventionRepository struct {
}

func NewSqliteInterventionRepository() *SqliteInterventionRepository {
	return &SqliteInterventionRepository{}
}

func (ir *SqliteInterventionRepository) AddIntervention(intervention models.Intervention) error {
	db, err := NewDatabase()
	if err != nil {
		return err
	}

	return db.AddIntervention(intervention)
}

func (ir *SqliteInterventionRepository) GetInterventions() ([]models.Intervention, error) {
	db, err := NewDatabase()
	if err != nil {
		return nil, err
	}

	return db.GetInterventions()
}

func (ir *SqliteInterventionRepository) GetIntervention(timestamp string) (models.Intervention, error) {
	db, err := NewDatabase()
	if err != nil {
		return models.Intervention{}, err
	}

	return db.GetIntervention(timestamp)
}

func (ir *SqliteInterventionRepository) UpdateIntervention(timestamp string, intervention models.Intervention) error {
	db, err := NewDatabase()
	if err != nil {
		return err
	}

	return db.UpdateIntervention(timestamp, intervention)
}
