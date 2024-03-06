package main

import (
	"log"

	"github.com/andreabertanzon/simtek-generate/data"
	"github.com/andreabertanzon/simtek-generate/generators"
)

const connectionString = "../simtek/simtek.db"

func main() {
	ir := data.NewSqliteInterventionRepository(connectionString)
	interventions, err := ir.GetInterventions()
	if err != nil {
		log.Fatal(err)
	}

	generator := generators.NewMarkdownGenerator()
	err = generator.Generate(interventions)
	if err != nil {
		log.Fatal(err)
	}
	// generator.Generate(interventions)
}
