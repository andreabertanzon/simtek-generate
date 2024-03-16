package main

import (
	"log"
	"os"

	"github.com/andreabertanzon/simtek-generate/data"
	"github.com/andreabertanzon/simtek-generate/generators"
)

const connectionString = "../simtek/simtek.db"

func main() {
	// get timestamp from command line
	timestamp := os.Args[1]

	ir := data.NewSqliteInterventionRepository(connectionString)
	/* timestamp := time.now().format("2006-01-02") */

	interventions, err := ir.GetInterventionsByDate(timestamp)

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
