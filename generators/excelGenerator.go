package generators

import (
	"strconv"
	"strings"
	"time"

	"github.com/andreabertanzon/simtek-generate/models"
	"github.com/xuri/excelize/v2"
)

type ExcelGenerator struct {
	Style Style
}

func NewExcelGenerator(*Style) *ExcelGenerator {
	return &ExcelGenerator{
		Style: Style{
			Font:            "Arial",
			Fill:            "solid",
			TextColor:       "black",
			FontH1Dimension: 20,
			FontH2Dimension: 18,
			FontH3Dimension: 16,
			PrimaryColor:    "#000000",
			SecondaryColor:  "#FFC0CB",
		},
	}
}

//TODO: Separate table to keep track of the generated files
//TODO: set the number of interventions generated in that file so that changes > replace/regenerate file

func (eg *ExcelGenerator) Generate(interventions []models.Intervention) error {
	today := time.Now().Format("02-01-2006")
	xlsx := excelize.NewFile()
	defer xlsx.Close()

	for interventionIndex, intervention := range interventions {
		currentRow := 1 + interventionIndex
		println(intervention.Site)
		sheetIndex, err := xlsx.NewSheet(intervention.Site)
		if err != nil {
			return err
		}
		// Site and Timestamp
		style, err := xlsx.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				Type:    "pattern",
				Color:   []string{"#FFC0CB"}, // Pink color
				Pattern: 1,                   // Solid pattern
			},
			Font: &excelize.Font{Bold: true},
		})
		if err != nil {
			return err
		}

		err = xlsx.SetCellStyle(intervention.Site, "A"+strconv.Itoa(currentRow), "B"+strconv.Itoa(currentRow), style)
		if err != nil {
			return err
		}

		xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), "Cantiere")
		xlsx.SetCellValue(intervention.Site, "B"+strconv.Itoa(currentRow), "Data")

		currentRow++

		xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), intervention.Site)
		xlsx.SetCellValue(intervention.Site, "B"+strconv.Itoa(currentRow), intervention.Timestamp)

		xlsx.SetActiveSheet(sheetIndex)

		// Workers
		currentRow += 2

		err = xlsx.SetCellStyle(intervention.Site, "A"+strconv.Itoa(currentRow), "B"+strconv.Itoa(currentRow), style)
		if err != nil {
			return err
		}

		xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), "Operatori")
		xlsx.SetCellValue(intervention.Site, "B"+strconv.Itoa(currentRow), "Ore")
		currentRow++

		for _, worker := range intervention.Workers {
			for workerName, workerHours := range worker {
				xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), workerName)
				xlsx.SetCellValue(intervention.Site, "B"+strconv.Itoa(currentRow), workerHours)
				currentRow++
			}
		}

		// Descriptions
		currentRow += 2

		err = xlsx.SetCellStyle(intervention.Site, "A"+strconv.Itoa(currentRow), "B"+strconv.Itoa(currentRow), style)
		if err != nil {
			return err
		}

		xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), "Descrizione Lavori")

		currentRow++

		for _, description := range intervention.Descriptions {
			xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), description)
			currentRow++
		}

		// Materials
		currentRow += 2

		err = xlsx.SetCellStyle(intervention.Site, "A"+strconv.Itoa(currentRow), "B"+strconv.Itoa(currentRow), style)
		if err != nil {
			return err
		}
		err = xlsx.SetCellStyle(intervention.Site, "C"+strconv.Itoa(currentRow), "C"+strconv.Itoa(currentRow), style)
		if err != nil {
			return err
		}

		xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), "Materiali")
		xlsx.SetCellValue(intervention.Site, "B"+strconv.Itoa(currentRow), "u.m.")
		xlsx.SetCellValue(intervention.Site, "C"+strconv.Itoa(currentRow), "Quantità")

		currentRow++

		for _, material := range intervention.Materials {
			materialParts := strings.Split(material, "|")
			xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), materialParts[0])
			xlsx.SetCellValue(intervention.Site, "B"+strconv.Itoa(currentRow), materialParts[1])
			xlsx.SetCellValue(intervention.Site, "C"+strconv.Itoa(currentRow), materialParts[2])
			currentRow++
		}

		// Notes
		currentRow += 2

		err = xlsx.SetCellStyle(intervention.Site, "A"+strconv.Itoa(currentRow), "B"+strconv.Itoa(currentRow), style)
		if err != nil {
			return err
		}

		xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), "Note")

		currentRow++

		xlsx.SetCellValue(intervention.Site, "A"+strconv.Itoa(currentRow), intervention.Notes)

		// Save the file
		err = xlsx.SaveAs(today + ".xlsx")
		if err != nil {
			return err
		}
	}
	return nil
}
