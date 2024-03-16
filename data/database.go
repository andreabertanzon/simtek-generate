package data

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/andreabertanzon/simtek-generate/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	ConnectionString string
}

func NewDatabase(connectionString string) (*Database, error) {
	obj := &Database{ConnectionString: connectionString}
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return obj, nil
}

func (d *Database) GetInterventions() ([]models.Intervention, error) {
	today := time.Now().Format("02-01-2006")

	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM interventions", today)
	if err != nil {
		return nil, err
	}

	var interventions []models.Intervention
	for rows.Next() {
		var id int
		var guid string
		var timestamp string
		var details string
		err = rows.Scan(&id, &guid, &timestamp, &details)
		if err != nil {
			return nil, err
		}
		intervention := models.Intervention{}
		err = json.Unmarshal([]byte(details), &intervention)
		if err != nil {
			return nil, err
		}

		interventions = append(interventions, intervention)

	}

	return interventions, nil
}

func (d *Database) GetInterventionsByDate(date string) ([]models.Intervention, error) {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM interventions WHERE timestamp = ?", date)
	if err != nil {
		return nil, err
	}

	var interventions []models.Intervention
	for rows.Next() {
		var id int
		var guid string
		var timestamp string
		var details string
		err = rows.Scan(&id, &guid, &timestamp, &details)
		if err != nil {
			return nil, err
		}
		intervention := models.Intervention{}
		err = json.Unmarshal([]byte(details), &intervention)
		if err != nil {
			return nil, err
		}

		interventions = append(interventions, intervention)

	}

	return interventions, nil
}

func (d *Database) GetIntervention(timestamp string) (models.Intervention, error) {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return models.Intervention{}, err
	}
	defer db.Close()

	var id int
	var guid string
	var details string
	selectSql := "SELECT * FROM interventions WHERE timestamp = ?"
	err = db.QueryRow(selectSql, timestamp).Scan(&id, &guid, &timestamp, &details)
	if err != nil {
		return models.Intervention{}, err
	}

	intervention := models.Intervention{}
	err = json.Unmarshal([]byte(details), &intervention)
	if err != nil {
		return models.Intervention{}, err
	}

	return intervention, nil
}

func (d *Database) UpdateIntervention(timestamp string, intervention models.Intervention) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	details, err := json.Marshal(intervention)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE interventions SET details = ? WHERE timestamp = ?", string(details), timestamp)
	if err != nil {
		return err
	}

	return nil
}
