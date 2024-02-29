package generators

import "github.com/andreabertanzon/simtek-generate/models"

type Generator interface {
	Generate([]models.Intervention) error
}
