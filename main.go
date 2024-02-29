package main

import (
	"log"

	"github.com/andreabertanzon/simtek-generate/data"
	"github.com/andreabertanzon/simtek-generate/generators"
)

func main() {
	ir := data.NewSqliteInterventionRepository()
	interventions, err := ir.GetInterventions()
	if err != nil {
		log.Fatal(err)
	}

	generator := generators.NewExcelGenerator(nil)
	generator.Generate(interventions)
}
