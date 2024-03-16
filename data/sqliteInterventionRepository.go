package data

import "github.com/andreabertanzon/simtek-generate/models"

type SqliteInterventionRepository struct {
	ConnectionString string
}

func NewSqliteInterventionRepository(connectionString string) *SqliteInterventionRepository {
	return &SqliteInterventionRepository{ConnectionString: connectionString}
}

func (ir *SqliteInterventionRepository) GetInterventions() ([]models.Intervention, error) {
	db, err := NewDatabase(ir.ConnectionString)
	if err != nil {
		return nil, err
	}

	return db.GetInterventions()
}

func (ir *SqliteInterventionRepository) GetInterventionsByDate(date string) ([]models.Intervention, error) {
	db, err := NewDatabase(ir.ConnectionString)
	if err != nil {
		return nil, err
	}

	return db.GetInterventionsByDate(date)
}

func (ir *SqliteInterventionRepository) GetIntervention(timestamp string) (models.Intervention, error) {
	db, err := NewDatabase(ir.ConnectionString)
	if err != nil {
		return models.Intervention{}, err
	}

	return db.GetIntervention(timestamp)
}
