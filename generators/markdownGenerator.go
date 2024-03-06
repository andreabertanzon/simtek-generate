package generators

import (
	"os"
	"text/template"

	"github.com/andreabertanzon/simtek-generate/models"
)

type MarkdownGenerator struct {
}

type WorkerMd struct {
	Name     string
	Hours    int
	HasSlash bool
}

func NewMarkdownGenerator() *MarkdownGenerator {
	return &MarkdownGenerator{}
}

// Generate a markdown file
func (md *MarkdownGenerator) Generate(interventions []models.Intervention) error {
	// Generating the title

	workers := []WorkerMd{{Name: "Simone", Hours: 4, HasSlash: false}, {Name: "Cristian", Hours: 4, HasSlash: true}}

	data := struct {
		Title           string
		Subtitle        string
		Workers         []WorkerMd
		LastWorkerIndex int
	}{
		Title:           "26-02-2024",
		Subtitle:        "Interventi del giorno",
		Workers:         workers,
		LastWorkerIndex: 2,
	}

	tmplSring, err := os.ReadFile("generated/test.md")
	if err != nil {
		return err
	}

	t, err := template.New("test").Parse(string(tmplSring))
	if err != nil {
		return err
	}

	file, err := os.Create("generated/test2.md")
	if err != nil {
		return err
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
