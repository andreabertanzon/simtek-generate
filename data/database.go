package data

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/andreabertanzon/simtek-generate/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
}

func NewDatabase() (*Database, error) {
	obj := &Database{}
	db, err := sql.Open("sqlite3", "simtek.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStmt := `
	create table if not exists interventions (
		id integer not null primary key autoincrement,
		timestamp text,
		details text
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (d *Database) AddIntervention(intervention models.Intervention) error {
	db, err := sql.Open("sqlite3", "simtek.db")
	if err != nil {
		return err
	}
	defer db.Close()

	details, err := json.Marshal(intervention)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO interventions (timestamp, details) VALUES (?, ?)", intervention.Timestamp, string(details))
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetInterventions() ([]models.Intervention, error) {
	today := time.Now().Format("02-01-2006")

	db, err := sql.Open("sqlite3", "simtek.db")
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
		var timestamp string
		var details string
		err = rows.Scan(&id, &timestamp, &details)
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
	db, err := sql.Open("sqlite3", "simtek.db")
	if err != nil {
		return models.Intervention{}, err
	}
	defer db.Close()

	var id int
	var details string
	err = db.QueryRow("SELECT * FROM interventions WHERE timestamp = ?", timestamp).Scan(&id, &timestamp, &details)
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
	db, err := sql.Open("sqlite3", "simtek.db")
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
