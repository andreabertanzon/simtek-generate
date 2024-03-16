package generators

import (
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/andreabertanzon/simtek-generate/models"
)

type MarkdownGenerator struct {
}

type WorkerMd struct {
	Name     string
	Hours    float32
	HasSlash bool
}

type Material struct {
	Material string
	Umeasure string
	Quantity float32
}

type MarkdownData struct {
	Title           string
	Subtitle        string
	Workers         []WorkerMd
	Materials       []Material
	Description     []string
	Notes           string
	LastWorkerIndex int
}

func NewMarkdownGenerator() *MarkdownGenerator {
	return &MarkdownGenerator{}
}

// Generate a markdown file
func (md *MarkdownGenerator) Generate(interventions []models.Intervention) error {
	// Generating the title
	// TODO: iterate over interventions and foreach intervention generate a markdown file
	// TODO: Once the files are generated (maybe in a concurrent way) merge them into a single file

	// Populate the workers array
	for _, intervention := range interventions {
		var data MarkdownData
		workers := []WorkerMd{}
		flattenedWorkers := models.Flatten(intervention.Workers)
		workerIndex := 0

		// ** WORKERS **
		for name, hour := range flattenedWorkers {
			hasSlash := false
			if workerIndex == len(flattenedWorkers)-1 {
				hasSlash = true
			}
			workerIndex++

			workers = append(workers, WorkerMd{Name: name, Hours: hour, HasSlash: hasSlash})
		}

		data.Workers = append(data.Workers, workers...)

		// ** LAST WORKER INDEX **
		data.LastWorkerIndex = len(data.Workers) - 1

		// ** TITLE **
		data.Title = intervention.Site
		data.Subtitle = "Intervento: " + intervention.Timestamp

		for _, description := range intervention.Descriptions {
			if description == " " || description == "" {
				continue
			}
			data.Description = append(data.Description, description)
		}

		for _, material := range intervention.Materials {
			splitMaterial := strings.Split(material, "|")

			if len(splitMaterial) != 3 {
				log.Println("Material not in the right format: ", material)
				continue
			}

			materialName := splitMaterial[0]
			umeasure := splitMaterial[1]
			quantity, err := strconv.ParseFloat(splitMaterial[2], 32)
			if err != nil {
				log.Printf("Error parsing quantity for material %s: %e", materialName, err)
				continue
			}

			material := Material{Material: materialName, Umeasure: umeasure, Quantity: float32(quantity)}
			data.Materials = append(data.Materials, material)
		}

		if intervention.Notes != "" {
			data.Notes = intervention.Notes
		} else {
			data.Notes = "Nessuna nota da segnalare."
		}

		tmplSring, err := os.ReadFile("templates/intervention.md")
		if err != nil {
			return err
		}

		t, err := template.New("test").Parse(string(tmplSring))
		if err != nil {
			return err
		}

		file, err := os.Create("generated/" + intervention.Site + ".md")
		if err != nil {
			return err
		}
		defer file.Close()

		err = t.Execute(file, data)
		if err != nil {
			return err
		}
	}

	return nil
}
