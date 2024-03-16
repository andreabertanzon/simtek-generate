package generators_test

import (
	"testing"

	"github.com/andreabertanzon/simtek-generate/generators"
	"github.com/andreabertanzon/simtek-generate/models"
	"github.com/google/uuid"
)

func TestMarkdownGenerator_Generate(t *testing.T) {
	interventions := []models.Intervention{
		{
			Guid:         uuid.New().String(),
			Site:         "Donini",
			Descriptions: []string{"This is a description"},
			Timestamp:    "2024-03-15",
			Workers: []map[string]int{
				{"John": 8},
				{"Jane": 8},
			},
		},
	}

	generator := generators.NewMarkdownGenerator()
	generator.Generate(interventions)
}
